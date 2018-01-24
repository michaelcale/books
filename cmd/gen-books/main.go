package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

var bookDirs = []string{
	"jQuery",
}

// KV represents a key/value pair
type KV struct {
	k string
	v string
}

// Section represents a part of a chapter
type Section struct {
	FilePath     string // path of the file from which we've read the section
	Title        string
	BodyMarkdown string
	data         []KV
}

// Chapter represents a book chapter
type Chapter struct {
	BookDir    string
	ChapterDir string
	IndexKV    []KV   // content of index.txt file
	Title      string // extracted from IndexKV
	Sections   []*Section
}

// Book represents a book
type Book struct {
	URL      string // used in index.tmpl.html
	Title    string // used in index.tmpl.html
	Chapters []*Chapter
}

func getV(a []KV, k string) (string, error) {
	for _, kv := range a {
		if kv.k == k {
			return kv.v, nil
		}
	}
	return "", fmt.Errorf("key '%s' not found", k)
}

func readFileAsLines(path string) ([]string, error) {
	d, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(d)
	res := strings.Split(s, "\n")
	return res, nil
}

func extractMultiLineValue(lines []string) ([]string, string, error) {
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "===" {
			rest := lines[i+1:]
			s := strings.Join(lines[:i], "\n")
			return rest, s, nil
		}
	}
	return nil, "", errors.New("didn't find end of value line ('===')")
}

// if error is io.EOF, we successfully finished parsing
func parseNextKV(lines []string) ([]string, KV, error) {
	// skip empty lines from the beginning
	var kv KV
	for len(lines) > 0 && len(lines[0]) == 0 {
		lines = lines[1:]
	}
	if len(lines) == 0 {
		return nil, kv, io.EOF
	}
	s := strings.TrimSpace(lines[0])
	lines = lines[1:]
	if !strings.HasSuffix(s, ":") {
		// this is singlie line "k: v"
		parts := strings.SplitN(s, ":", 2)
		if len(parts) != 2 {
			return nil, kv, fmt.Errorf("'%s' is not a valid start for k/v", s)
		}
		kv.k, kv.v = parts[0], parts[1]
		return lines, kv, nil
	}
	// this is a multi-line value that ends with '==='
	kv.k = strings.TrimSuffix(s, ":")
	var err error
	lines, kv.v, err = extractMultiLineValue(lines)
	return lines, kv, err
}

/*
parseKVFile parsers my brand of key/value text file optimized for human editing
Key/value are encoded in 2 ways:
1. On a single line, if value is short and doesn't contain '\n'

"key: value\n"

2. On multiple lines, if value is long or contains '\n'

key:
value
===\n
*/
func parseKVFile(path string) ([]KV, error) {
	lines, err := readFileAsLines(path)
	var res []KV
	var kv KV
	for {
		lines, kv, err = parseNextKV(lines)
		if err == io.EOF {
			return res, nil
		}
		if err != nil {
			return nil, err
		}
		res = append(res, kv)
	}
}

func shortenV(v string) string {
	parts := strings.SplitN(v, "\n", 2)
	s := parts[0]
	if len(s) < 60 {
		return s
	}
	return s[:60] + "..."
}

func dumpKV(a []KV) {
	for _, kv := range a {
		fmt.Printf("K: %s\nV: %s\n", kv.k, shortenV(kv.v))
	}
}

func parseSection(path string) (*Section, error) {
	kv, err := parseKVFile(path)
	if err != nil {
		return nil, err
	}
	res := &Section{
		FilePath: path,
		data:     kv,
	}
	res.Title, err = getV(kv, "Title")
	if err != nil {
		return nil, err
	}
	res.BodyMarkdown, err = getV(kv, "Body")
	if err != nil {
		dumpKV(kv)
		err = fmt.Errorf("parseSection('%s'), err: '%s'", path, err)
		return nil, err
	}
	return res, nil
}

func parseChapter(chapter *Chapter) error {
	dir := filepath.Join(chapter.BookDir, chapter.ChapterDir)
	path := filepath.Join(dir, "index.txt")
	indexKV, err := parseKVFile(path)
	if err != nil {
		return err
	}
	chapter.IndexKV = indexKV
	fileInfos, err := ioutil.ReadDir(dir)
	for _, fi := range fileInfos {
		if fi.IsDir() || !fi.Mode().IsRegular() {
			continue
		}
		name := fi.Name()
		if strings.ToLower(filepath.Ext(name)) != ".md" {
			continue
		}
		path = filepath.Join(dir, name)
		section, err := parseSection(path)
		if err != nil {
			return err
		}
		chapter.Sections = append(chapter.Sections, section)
	}
	return nil
}

func genBook(bookName string) (*Book, error) {
	bookDirName := makeURLSafe(bookName)
	book := &Book{
		Title: bookName,
		URL:   fmt.Sprintf("/book/%s/", bookDirName),
	}
	bookDir := filepath.Join("book", bookDirName)
	fileInfos, err := ioutil.ReadDir(bookDir)
	if err != nil {
		return nil, err
	}
	var chapters []*Chapter
	for _, fi := range fileInfos {
		if fi.IsDir() {
			ch := &Chapter{
				BookDir:    bookDir,
				ChapterDir: fi.Name(),
			}
			err = parseChapter(ch)
			if err != nil {
				return nil, err
			}
			chapters = append(chapters, ch)
			continue
		}
		return nil, fmt.Errorf("Unexpected file at top-level: '%s'", fi.Name())
	}
	nSections := 0
	for _, ch := range chapters {
		nSections += len(ch.Sections)
	}
	fmt.Printf("Book '%s' %d chapters, %d sections\n", bookName, len(chapters), nSections)
	book.Chapters = chapters
	return book, nil
}

func main() {
	var books []*Book
	for _, bookName := range bookDirs {
		timeStart := time.Now()
		book, err := genBook(bookName)
		if err != nil {
			fmt.Printf("Error '%s' parsing book '%s'\n", err, bookName)
			return
		}
		books = append(books, book)
		fmt.Printf("Generating book '%s' took %s\n", bookName, time.Since(timeStart))
	}
	genIndex(books)
}
