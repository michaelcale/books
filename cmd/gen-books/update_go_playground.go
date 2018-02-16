package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kjk/u"
)

func calcSourceFileHashesInner(dir string, fileToSha1Chan chan map[string]string) {
	timeStart := time.Now()
	m := make(map[string]string)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".go" {
			return nil
		}
		sha1Hex, err := u.Sha1HexOfFile(path)
		u.PanicIfErr(err)
		m[path] = sha1Hex
		return nil
	})
	fmt.Printf("calcSourceFileHashesInner() took %s\n", time.Since(timeStart))
	fileToSha1Chan <- m
}

// calculate sha1 of source .go files on a separate goroutine
// to speed up things
func calcSourceFileHashes(dir string) chan map[string]string {
	ch := make(chan map[string]string)
	go calcSourceFileHashesInner(dir, ch)
	return ch
}

// EmbeddedGoSourceFile describes embedded source file as described
// by @file line in markdown sources
type EmbeddedGoSourceFile struct {
	FileName          string
	Sha1Hex           string
	PlaygroundShareID string
	// index into MarkdownFile.Lines
	LineNo int
}

// GoMarkdownFile describes a single markdown file
type GoMarkdownFile struct {
	Path                string
	Data                []byte
	Lines               []string
	EmbeddedSourceFiles []EmbeddedGoSourceFile
}

func dataToLines(d []byte) []string {
	s := string(d)
	return strings.Split(s, "\n")
}

func extractEmbeddedGoSourceFiles(lines []string) []EmbeddedGoSourceFile {
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
		f := EmbeddedGoSourceFile{
			FileName:          fileDirective.FileName,
			Sha1Hex:           fileDirective.Sha1Hex,
			PlaygroundShareID: fileDirective.PlaygroundID,
			LineNo:            i,
		}
		res = append(res, f)
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
		embeddedFiles := extractEmbeddedGoSourceFiles(lines)
		if len(embeddedFiles) == 0 {
			return nil
		}
		f := GoMarkdownFile{
			Path:                path,
			Lines:               lines,
			Data:                d,
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
	return "", nil
	// TODO: write me
}

// // ${dir} is books/go/
func updateGoPlaygroundLinks(dir string) {
	fmt.Printf("updateGoPlaygroundLinks() started\n")
	timeStart := time.Now()
	fileToSha1Chan := calcSourceFileHashes(dir)
	markdownFiles := loadMarkdownFiles(dir)
	fileToSha1 := <-fileToSha1Chan
	for _, mf := range markdownFiles {
		for _, ef := range mf.EmbeddedSourceFiles {
			fullFileName := filepath.Join(filepath.Dir(mf.Path), ef.FileName)
			realSha1 := fileToSha1[fullFileName]
			u.PanicIf(realSha1 == "", "didn't find sha1 for file '%s' ('%s') in markdown file '%s'", fullFileName, ef.FileName, mf.Path)
			if ef.PlaygroundShareID != "" && (ef.Sha1Hex == realSha1) {
				// skipping if we already have playground id and the file didn't change since the time we
				continue
			}
			fmt.Printf("Getting playground share id for '%s'\n", fullFileName)
			_, err := getGoPlaygroundShareID(mf.Data)
			u.PanicIfErr(err)
		}
	}
	fmt.Printf("updateGoPlaygroundLinks() finished in %s\n", time.Since(timeStart))
}
