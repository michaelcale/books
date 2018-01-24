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
	chapter        *Chapter
	SourceFilePath string // path of the file from which we've read the section
	Title          string // used in book_index.tmpl.html
	TitleSafe      string
	BodyMarkdown   string
	data           []KV
}

// URL returns url of .html file with this section
func (s *Section) URL() string {
	chap := s.chapter
	book := chap.book
	bookTitle := book.TitleSafe
	chapTitle := chap.TitleSafe
	sectionTitle := s.TitleSafe
	return fmt.Sprintf("/book/%s/%s/%s.html", bookTitle, chapTitle, sectionTitle)
}

func (s *Section) destFilePath() string {
	chap := s.chapter
	book := chap.book
	bookTitle := book.TitleSafe
	chapTitle := chap.TitleSafe
	sectionTitle := s.TitleSafe + ".html"
	return filepath.Join("book_html", bookTitle, chapTitle, sectionTitle)
}

// Chapter represents a book chapter
type Chapter struct {
	book       *Book
	ChapterDir string
	IndexKV    []KV   // content of index.txt file
	Title      string // extracted from IndexKV, used in book_index.tmpl.html
	TitleSafe  string
	Sections   []*Section
}

// URL is used in book_index.tmpl.html
func (c *Chapter) URL() string {
	book := c.book
	bookTitle := book.TitleSafe
	chapTitle := c.TitleSafe
	return fmt.Sprintf("/book/%s/%s/index.html", bookTitle, chapTitle)
}

func (c *Chapter) destFilePath() string {
	book := c.book
	bookTitle := book.TitleSafe
	chapTitle := c.TitleSafe
	return filepath.Join("book_html", bookTitle, chapTitle, "index.html")
}

// Book represents a book
type Book struct {
	URL       string // used in index.tmpl.html
	Title     string // used in index.tmpl.html
	TitleLong string // used in book_index.tmpl.html
	TitleSafe string
	Chapters  []*Chapter
	SourceDir string // dir where source markdown files are
	DestDir   string // dif where destitation html files are
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
		SourceFilePath: path,
		data:           kv,
	}
	res.Title, err = getV(kv, "Title")
	if err != nil {
		return nil, err
	}
	res.TitleSafe = makeURLSafe(res.Title)
	res.BodyMarkdown, err = getV(kv, "Body")
	if err != nil {
		dumpKV(kv)
		err = fmt.Errorf("parseSection('%s'), err: '%s'", path, err)
		return nil, err
	}
	return res, nil
}

func parseChapter(chapter *Chapter) error {
	dir := filepath.Join(chapter.book.SourceDir, chapter.ChapterDir)
	path := filepath.Join(dir, "index.txt")
	indexKV, err := parseKVFile(path)
	if err != nil {
		return err
	}
	chapter.IndexKV = indexKV
	chapter.Title, err = getV(indexKV, "Title")
	chapter.TitleSafe = makeURLSafe(chapter.Title)
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
		section.chapter = chapter
		chapter.Sections = append(chapter.Sections, section)
	}
	return nil
}

func parseBook(bookName string) (*Book, error) {
	bookNameSafe := makeURLSafe(bookName)
	dir := filepath.Join("book", bookNameSafe)
	book := &Book{
		Title:     bookName,
		TitleLong: fmt.Sprintf("Essential %s", bookName),
		TitleSafe: bookNameSafe,
		SourceDir: dir,
		DestDir:   filepath.Join("book_html", "book", bookNameSafe),
		URL:       fmt.Sprintf("/book/%s/", bookNameSafe),
	}
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var chapters []*Chapter
	for _, fi := range fileInfos {
		if fi.IsDir() {
			ch := &Chapter{
				book:       book,
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
		book, err := parseBook(bookName)
		if err != nil {
			fmt.Printf("Error '%s' parsing book '%s'\n", err, bookName)
			return
		}
		books = append(books, book)
		fmt.Printf("Generating book '%s' took %s\n", bookName, time.Since(timeStart))
	}
	genIndex(books)
	for _, book := range books {
		genBook(book)
	}
}
