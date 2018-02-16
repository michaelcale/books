package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kjk/u"
)

// EmbeddedGoSourceFile describes embedded source file as described
// by @file line in markdown sources
type EmbeddedGoSourceFile struct {
	FileName      string
	Path          string
	FileDirective *FileDirective

	// content of the file after filtering
	Lines         []string
	cachedData    []byte
	cachedSha1Hex string

	// index into MarkdownFile.Lines
	LineNo int
}

// Data returns content of the file
func (f *EmbeddedGoSourceFile) Data() []byte {
	if len(f.cachedData) == 0 {
		s := strings.Join(f.Lines, "\n")
		f.cachedData = []byte(s)
	}
	return f.cachedData
}

// RealSha1Hex returns hex version of sha1 of file content
func (f *EmbeddedGoSourceFile) RealSha1Hex() string {
	if f.cachedSha1Hex == "" {
		f.cachedSha1Hex = u.Sha1HexOfBytes(f.Data())
	}
	return f.cachedSha1Hex
}

// GoMarkdownFile describes a single markdown file
type GoMarkdownFile struct {
	Path                string
	Lines               []string
	EmbeddedSourceFiles []EmbeddedGoSourceFile
}

// Data returns files content reconstructed from Lines
func (f *GoMarkdownFile) Data() []byte {
	s := strings.Join(f.Lines, "\n")
	return []byte(s)
}

func dataToLines(d []byte) []string {
	s := string(d)
	return strings.Split(s, "\n")
}

func readFilteredSourceFile(path string) ([]string, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := dataToLines(d)
	lines = removeAnnotationLines(lines)
	return lines, nil
}

func extractEmbeddedGoSourceFiles(mdPath string, lines []string) []EmbeddedGoSourceFile {
	dir := filepath.Dir(mdPath)
	var res []EmbeddedGoSourceFile
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "@file ") {
			continue
		}
		fileDirective, err := parseFileDirective(line)
		u.PanicIfErr(err)
		if fileDirective.NoPlayground {
			continue
		}
		filePath := filepath.Join(dir, fileDirective.FileName)
		lines, err := readFilteredSourceFile(filePath)
		u.PanicIfErr(err)
		f := EmbeddedGoSourceFile{
			FileName:      fileDirective.FileName,
			Path:          filePath,
			FileDirective: fileDirective,
			Lines:         lines,
			LineNo:        i,
		}
		res = append(res, f)
	}
	return res
}

// we don't want to show our // :show annotations in snippets
func removeAnnotationLines(lines []string) []string {
	var res []string
	prevWasEmpty := false
	for _, l := range lines {
		if strings.Contains(l, "// :show ") {
			continue
		}
		if len(l) == 0 && prevWasEmpty {
			continue
		}
		prevWasEmpty = len(l) == 0
		res = append(res, l)
	}
	return res
}

func loadMarkdownFiles(dir string) []GoMarkdownFile {
	timeStart := time.Now()
	var res []GoMarkdownFile
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".md" {
			return nil
		}
		d, err := ioutil.ReadFile(path)
		u.PanicIfErr(err)
		if !bytes.Contains(d, []byte("@file ")) {
			return nil
		}
		lines := dataToLines(d)
		embeddedFiles := extractEmbeddedGoSourceFiles(path, lines)
		if len(embeddedFiles) == 0 {
			return nil
		}
		f := GoMarkdownFile{
			Path:                path,
			Lines:               lines,
			EmbeddedSourceFiles: embeddedFiles,
		}
		res = append(res, f)
		return nil
	})
	fmt.Printf("loadMarkdownFiles took %s\n", time.Since(timeStart))
	return res
}

// submit the data to Go playground and get share id
func getGoPlaygroundShareID(d []byte) (string, error) {
	uri := "https://play.golang.org/share"
	r := bytes.NewBuffer(d)
	resp, err := http.Post(uri, "text/plain", r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("http.Post returned error code '%s'", err)
	}
	d, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(d)), nil
}

func testGetGoPlaygroundShareIDAndExit() {
	path := "books/go/0230-mutex/rwlock.go"
	d, err := ioutil.ReadFile(path)
	u.PanicIfErr(err)
	shareID, err := getGoPlaygroundShareID(d)
	u.PanicIfErr(err)
	fmt.Printf("share id: '%s'\n", shareID)
	os.Exit(0)
}

// // ${dir} is books/go/
func updateGoPlaygroundLinks(dir string) {
	fmt.Printf("updateGoPlaygroundLinks() started\n")
	timeStart := time.Now()
	markdownFiles := loadMarkdownFiles(dir)
	max := 10
	for _, mf := range markdownFiles {
		wasChanged := false
		for _, ef := range mf.EmbeddedSourceFiles {
			fullFileName := filepath.Join(filepath.Dir(mf.Path), ef.FileName)
			realSha1 := ef.RealSha1Hex()
			u.PanicIf(realSha1 == "", "didn't find sha1 for file '%s' ('%s') in markdown file '%s'", fullFileName, ef.FileName, mf.Path)
			if ef.FileDirective.GoPlaygroundID != "" && (ef.FileDirective.Sha1Hex == realSha1) {
				// skipping if we already have playground id and the file didn't change since the time we
				continue
			}
			fmt.Printf("Getting playground share id for '%s'\n", fullFileName)
			shareID, err := getGoPlaygroundShareID(ef.Data())
			u.PanicIfErr(err)
			ef.FileDirective.Sha1Hex = realSha1
			ef.FileDirective.GoPlaygroundID = shareID
			mf.Lines[ef.LineNo] = ef.FileDirective.String()
			wasChanged = true
		}
		if wasChanged {
			err := ioutil.WriteFile(mf.Path, mf.Data(), 0644)
			u.PanicIfErr(err)
			fmt.Printf("Wrote updated '%s'\n", mf.Path)
			max--
		}
		if max <= 0 {
			goto Exit
		}
	}
Exit:
	fmt.Printf("updateGoPlaygroundLinks() finished in %s\n", time.Since(timeStart))
}
