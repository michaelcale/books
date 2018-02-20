package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kjk/u"
)

const (
	showStartLine = "// :show start"
	showEndLine   = "// :show end"
)

func isShowStart(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return s == showStartLine
}

func isShowEnd(s string) bool {
	s = strings.ToLower(strings.TrimSpace(s))
	return s == showEndLine
}

func countStartChars(s string, c byte) int {
	for i := range s {
		if s[i] != c {
			return i
		}
	}
	return len(s)
}

// remove longest common space/tab prefix on non-empty lines
func shiftLines(lines []string) {
	maxTabPrefix := 1024
	maxSpacePrefix := 1024
	// first determine how much we can remove
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		n := countStartChars(line, ' ')
		if n > 0 {
			if n < maxSpacePrefix {
				maxSpacePrefix = n
			}
			continue
		}
		n = countStartChars(line, '\t')
		if n > 0 {
			if n < maxTabPrefix {
				maxTabPrefix = n
			}
			continue
		}
		// if doesn't start with space or tab, early abort
		return
	}
	if maxSpacePrefix == 1024 && maxTabPrefix == 1024 {
		return
	}

	toRemove := maxSpacePrefix
	if maxTabPrefix != 1024 {
		toRemove = maxTabPrefix
	}
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		lines[i] = line[toRemove:]
	}
}

// removes empty lines from the beginning and end of the array
func trimEmptyLines(lines []string) []string {
	for len(lines) > 0 && len(lines[0]) == 0 {
		lines = lines[1:]
	}
	for len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func extractCodeSnippets(path string) ([]string, error) {
	fc, err := loadFileCached(path)
	if err != nil {
		return nil, err
	}
	lines := fc.Lines
	nLines := len(lines)
	res := make([]string, 0, nLines)
	inShow := false
	for _, line := range lines {
		if isShowStart(line) {
			if inShow {
				return nil, fmt.Errorf("file '%s': consequitive '%s' lines", path, showStartLine)
			}
			inShow = true
			continue
		}
		if isShowEnd(line) {
			if !inShow {
				return nil, fmt.Errorf("file '%s': '%s' without start line", path, showEndLine)
			}
			inShow = false
			// add a separation line between show sections.
			// should be the right thing more often than not
			res = append(res, "")
			continue
		}
		if inShow {
			res = append(res, line)
		}
	}
	// if there are no show: markings, assume we want to show the whole file
	if len(res) == 0 {
		return trimEmptyLines(lines), nil
	}
	shiftLines(res)
	return trimEmptyLines(res), nil
}

func getLangFromFileExt(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".go":
		return "go"
	}
	fmt.Printf("Couldn't deduce language from file name '%s'\n", fileName)
	// TODO: more languages
	return ""
}

// replace potentially windows paths \foo\bar into unix paths /foo/bar
func toUnixPath(s string) string {
	return strings.Replace(s, `\`, "/", -1)
}

// convert local path like books/go/foo.go into path to the file in a github repo
func getGitHubPathForFile(path string) string {
	return "https://github.com/essentialbooks/books/blob/master/" + toUnixPath(path)
}

// FileDirective describes result of parsing
// @file ${fileName} output allow_error
type FileDirective struct {
	FileName       string
	WithOutput     bool
	AllowError     bool
	NoPlayground   bool
	Sha1Hex        string
	GoPlaygroundID string
}

// String serializes FileDirective back to string format
func (fd *FileDirective) String() string {
	s := fmt.Sprintf("@file %s", fd.FileName)
	if fd.WithOutput {
		s += " output"
	}
	if fd.AllowError {
		s += " allow_error"
	}
	if fd.NoPlayground {
		return s + " no_playground"
	}
	if fd.Sha1Hex != "" {
		s += " sha1:" + fd.Sha1Hex
	}
	if fd.GoPlaygroundID != "" {
		s += " goplayground:" + fd.GoPlaygroundID
	}
	return s
}

// parseFileDirective parses line like:
// @file ${fileName} [output] [allow_error] [no_playground] [noplayground] [sha1:${sha1}] [goplayground:${playgroundID}]
// into FileDirective
func parseFileDirective(line string) (*FileDirective, error) {
	line = strings.TrimSpace(line)
	u.PanicIf(!strings.HasPrefix(line, "@file"))
	parts := strings.Split(line, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid @file line: '%s'", line)
	}
	if parts[0] != "@file" {
		return nil, fmt.Errorf("invalid @file line: '%s'", line)
	}
	res := &FileDirective{}
	parts = parts[1:]
	res.FileName = parts[0]
	parts = parts[1:]
	for _, s := range parts {
		if len(s) == 0 {
			continue
		}
		s = strings.TrimSpace(s)
		switch {
		case s == "output":
			res.WithOutput = true
		case s == "allow_error":
			res.AllowError = true
		case s == "no_playground" || s == "noplayground":
			res.NoPlayground = true
		case strings.HasPrefix(s, "sha1:"):
			parts := strings.Split(s, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid sha1 in '%s'", line)
			}
			sha1Hex := parts[1]
			// 20 bytes * 2 bytes per char in hex encoding
			if len(sha1Hex) != 40 {
				return nil, fmt.Errorf("invalid sha1: in '%s'", line)
			}
			res.Sha1Hex = sha1Hex
		case strings.HasPrefix(s, "goplayground:"):
			parts := strings.Split(s, ":")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid playground: in '%s'", line)
			}
			res.GoPlaygroundID = parts[1]
		default:
			return nil, fmt.Errorf("invalid @file line: '%s', unknown option '%s'", line, s)
		}
	}
	return res, nil
}

// ${baseDir} is books/go/
// loads a source file whose name is in ${line} and
func extractCodeSnippetsAsMarkdownLines(baseDir string, line string) ([]string, error) {
	// line is:
	// @file ${fileName} [output]
	directive, err := parseFileDirective(line)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(baseDir, directive.FileName)
	if !fileExists(path) {
		return nil, fmt.Errorf("no file '%s' in line '%s'", path, line)
	}
	lines, err := extractCodeSnippets(path)
	if err != nil {
		return nil, err
	}
	lang := getLangFromFileExt(path)
	sep := "|"
	u.PanicIf(strings.Contains(lang, sep), "lang ('%s') contains '%s'", lang, sep)
	u.PanicIf(strings.Contains(path, sep), "path ('%s') contains '%s'", path, sep)
	// this line is parsed in parseCodeBlockInfo
	s := fmt.Sprintf("%s|github|%s", lang, getGitHubPathForFile(path))
	if directive.GoPlaygroundID != "" {
		// alternative would be https://play.golang.org/p/ + ${id}
		uri := "https://goplay.space/#" + directive.GoPlaygroundID
		s += "|playground|" + uri
	}
	res := []string{"```" + s}
	res = append(res, lines...)
	res = append(res, "```")

	if !directive.WithOutput {
		return res, nil
	}

	out, err := getOutput(path)
	if err != nil && !directive.AllowError {
		fmt.Printf("getOutput('%s'): error '%s', output: '%s'\n", path, err, out)
		maybePanicIfErr(err)
		return res, err
	}

	res = append(res, "")
	res = append(res, "**Output**:")
	res = append(res, "")
	res = append(res, "```text")
	lines = strings.Split(out, "\n")
	lines = trimEmptyLines(lines)
	res = append(res, lines...)
	res = append(res, "```")
	return res, nil
}

// runs `go run ${path}` and returns captured output`
func getGoOutput(path string) (string, error) {
	cmd := exec.Command("go", "run", path)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// it executes a code file and captures the output
func getOutput(path string) (string, error) {
	ext := strings.ToLower(filepath.Ext(path))

	if ext == ".go" {
		return getGoOutput(path)
	}
	return "", fmt.Errorf("getOutpu(%s): files with extension '%s' are not supported", path, ext)
}
