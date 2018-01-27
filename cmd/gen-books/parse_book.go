package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kjk/programming-books/pkg/common"
	"github.com/kjk/programming-books/pkg/kvstore"
	"github.com/kjk/u"
)

// SectionRef describes a sibling section
type SectionRef struct {
	IsCurrent bool
	Title     string
	URL       string
	No        int
}

// Section represents a part of a chapter
type Section struct {
	Chapter        *Chapter
	SourceFilePath string // path of the file from which we've read the section
	Title          string // used in book_index.tmpl.html
	TitleSafe      string
	BodyMarkdown   string
	BodyHTML       template.HTML
	No             int
	SectionRefs    []SectionRef
	doc            kvstore.Doc
}

// Book retuns book this section belongs to
func (s *Section) Book() *Book {
	return s.Chapter.Book
}

// GitHubText returns text we display in GitHub box
func (s *Section) GitHubText() string {
	return "Edit on GitHub"
}

// GitHubURL returns url to GitHub repo
func (s *Section) GitHubURL() string {
	uri := s.Chapter.GitHubURL() + "/" + filepath.Base(s.SourceFilePath)
	uri = strings.Replace(uri, "/tree/", "/blob/", -1)
	return uri
}

// URL returns url of .html file with this section
func (s *Section) URL() string {
	chap := s.Chapter
	book := chap.Book
	bookTitle := book.TitleSafe
	chapTitle := chap.TitleSafe
	sectionTitle := s.TitleSafe
	return fmt.Sprintf("/book/%s/%s/%s", bookTitle, chapTitle, sectionTitle)
}

func (s *Section) destFilePath() string {
	chap := s.Chapter
	book := chap.Book
	bookTitle := book.TitleSafe
	chapTitle := chap.TitleSafe
	sectionTitle := s.TitleSafe + ".html"
	return filepath.Join("books_html", "book", bookTitle, chapTitle, sectionTitle)
}

// Chapter represents a book chapter
type Chapter struct {
	Book       *Book
	ChapterDir string
	IndexDoc   kvstore.Doc // content of index.txt file
	Title      string      // extracted from IndexKV, used in book_index.tmpl.html
	TitleSafe  string
	Sections   []*Section
	No         int
	IsCurrent  bool
}

// GitHubText returns text we display in GitHub box
func (c *Chapter) GitHubText() string {
	return "Edit on GitHub"
}

// GitHubURL returns url to GitHub repo
func (c *Chapter) GitHubURL() string {
	return c.Book.GitHubURL() + "/" + c.ChapterDir
}

// VersionsHTML returns html version of versions
func (c *Chapter) VersionsHTML() template.HTML {
	s, err := c.IndexDoc.GetV("VersionsHtml")
	if err != nil {
		s = ""
	}
	return template.HTML(s)
}

// URL is used in book_index.tmpl.html
func (c *Chapter) URL() string {
	book := c.Book
	bookTitle := book.TitleSafe
	chapTitle := c.TitleSafe
	return fmt.Sprintf("/book/%s/%s/", bookTitle, chapTitle)
}

func (c *Chapter) destFilePath() string {
	book := c.Book
	bookTitle := book.TitleSafe
	chapTitle := c.TitleSafe
	return filepath.Join("books_html", "book", bookTitle, chapTitle, "index.html")
}

// IntroductionHTML retruns html version of Introduction:
func (c *Chapter) IntroductionHTML() template.HTML {
	s, err := c.IndexDoc.GetV("Introduction")
	if err != nil {
		return template.HTML("")
	}
	html := markdownToHTML([]byte(s), "")
	return template.HTML(html)
}

// SyntaxHTML retruns html version of Syntax:
func (c *Chapter) SyntaxHTML() template.HTML {
	s, err := c.IndexDoc.GetV("Syntax")
	if err != nil {
		return template.HTML("")
	}
	html := markdownToHTML([]byte(s), "")
	return template.HTML(html)
}

// RemarksHTML retruns html version of Remarks:
func (c *Chapter) RemarksHTML() template.HTML {
	s, err := c.IndexDoc.GetV("Remarks")
	if err != nil {
		return template.HTML("")
	}
	html := markdownToHTML([]byte(s), "")
	return template.HTML(html)
}

// ContributorsHTML retruns html version of Contributors:
func (c *Chapter) ContributorsHTML() template.HTML {
	s, err := c.IndexDoc.GetV("Contributors")
	if err != nil {
		return template.HTML("")
	}
	html := markdownToHTML([]byte(s), "")
	return template.HTML(html)
}

// Book represents a book
type Book struct {
	URL            string // used in index.tmpl.html
	Title          string // used in index.tmpl.html
	TitleLong      string // used in book_index.tmpl.html
	TitleSafe      string
	Chapters       []*Chapter
	SourceDir      string // dir where source markdown files are
	DestDir        string // dif where destitation html files are
	SoContributors []int

	cachedSectionsCount int
	defaultLang         string // default programming language for programming examples
}

// GitHubText returns text we show in GitHub link
func (b *Book) GitHubText() string {
	return "Edit on GitHub"
}

// GitHubURL returns link to GitHub for this book
func (b *Book) GitHubURL() string {
	return gitHubBaseURL + "/tree/master/books/" + filepath.Base(b.DestDir)
}

// SectionsCount returns total number of sections
func (b *Book) SectionsCount() int {
	if b.cachedSectionsCount != 0 {
		return b.cachedSectionsCount
	}
	nSections := 0
	for _, ch := range b.Chapters {
		nSections += len(ch.Sections)
	}
	b.cachedSectionsCount = nSections
	return nSections
}

// ChaptersCount returns number of chapters
func (b *Book) ChaptersCount() int {
	return len(b.Chapters)
}

var (
	defTitle = "No Title"
)

func dumpKV(doc kvstore.Doc) {
	for _, kv := range doc {
		fmt.Printf("K: %s\nV: %s\n", kv.Key, common.ShortenString(kv.Value))
	}
}

func parseSection(path string) (*Section, error) {
	doc, err := kvstore.ParseKVFile(path)
	if err != nil {
		fmt.Printf("Error parsing KV file: '%s'\n", path)
		return nil, err
	}
	res := &Section{
		SourceFilePath: path,
		doc:            doc,
	}
	res.Title = doc.GetVSilent("Title", defTitle)
	if res.Title == defTitle {
		fmt.Printf("parseSection: no title for %s\n", path)
	}
	res.TitleSafe = common.MakeURLSafe(res.Title)
	res.BodyMarkdown, err = doc.GetV("Body")
	if err == nil {
		return res, nil
	}
	s, err := doc.GetV("BodyHtml")
	res.BodyHTML = template.HTML(s)
	if err == nil {
		return res, nil
	}
	dumpKV(doc)
	return nil, fmt.Errorf("parseSection('%s'), err: '%s'", path, err)
}

func refSectionSetCurrent(refs []SectionRef, activeNo int) []SectionRef {
	var res []SectionRef
	for i, ref := range refs {
		ref.IsCurrent = (i == activeNo)
		res = append(res, ref)
	}
	return res
}

func buildSectionRefs(sections []*Section) {
	var refs []SectionRef
	for i, section := range sections {
		ref := SectionRef{
			IsCurrent: false,
			Title:     section.Title,
			URL:       section.URL(),
			No:        i + 1,
		}
		refs = append(refs, ref)
	}
	for i, section := range sections {
		section.SectionRefs = refSectionSetCurrent(refs, i)
	}
}

func parseChapter(chapter *Chapter) error {
	dir := filepath.Join(chapter.Book.SourceDir, chapter.ChapterDir)
	path := filepath.Join(dir, "index.txt")
	doc, err := kvstore.ParseKVFile(path)
	if err != nil {
		return err
	}
	chapter.IndexDoc = doc
	chapter.Title, err = doc.GetV("Title")
	chapter.TitleSafe = common.MakeURLSafe(chapter.Title)
	fileInfos, err := ioutil.ReadDir(dir)
	var sections []*Section
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
		section.Chapter = chapter
		section.No = len(sections) + 1
		sections = append(sections, section)
	}
	buildSectionRefs(sections)
	chapter.Sections = sections
	return nil
}

func soContributorURL(userID int) string {
	return fmt.Sprintf("https://stackoverflow.com/users/%d/", userID)
}

func loadSoContributorsMust(book *Book, path string) {
	lines, err := common.ReadFileAsLines(path)
	u.PanicIfErr(err)
	var ids []int
	for _, line := range lines {
		id, err := strconv.Atoi(line)
		u.PanicIfErr(err)
		ids = append(ids, id)
	}
	book.SoContributors = ids
}

// TODO: add github contributors
func genContributorsMarkdown(soUserIDs []int) string {
	if len(soUserIDs) == 0 {
		return ""
	}
	lines := []string{
		"Contributors from Stack Overflow:",
	}
	for _, userID := range soUserIDs {
		s := fmt.Sprintf("* [%d](%s)", userID, soContributorURL(userID))
		lines = append(lines, s)
	}
	return strings.Join(lines, "\n")
}

func genContributorsChapter(book *Book) *Chapter {
	md := genContributorsMarkdown(book.SoContributors)
	var doc kvstore.Doc
	kv := kvstore.KeyValue{
		Key:   "Contributors",
		Value: md,
	}
	doc = append(doc, kv)
	ch := &Chapter{
		Book:      book,
		IndexDoc:  doc,
		Title:     "Contributors",
		TitleSafe: common.MakeURLSafe("Contributors"),
		No:        999,
	}
	return ch
}

func parseBook(bookName string) (*Book, error) {
	bookNameSafe := common.MakeURLSafe(bookName)
	dir := filepath.Join("books", bookNameSafe)
	book := &Book{
		Title:     bookName,
		TitleLong: fmt.Sprintf("Essential %s", bookName),
		TitleSafe: bookNameSafe,
		SourceDir: dir,
		DestDir:   filepath.Join("books_html", "book", bookNameSafe),
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
				Book:       book,
				ChapterDir: fi.Name(),
			}
			err = parseChapter(ch)
			if err != nil {
				return nil, err
			}
			chapters = append(chapters, ch)
			ch.No = len(chapters)
			continue
		}

		if fi.Name() == "so_contributors.txt" {
			path := filepath.Join(dir, fi.Name())
			loadSoContributorsMust(book, path)
			continue
		}
		return nil, fmt.Errorf("Unexpected file at top-level: '%s'", fi.Name())
	}
	ch := genContributorsChapter(book)
	ch.No = len(chapters) + 1
	chapters = append(chapters, ch)
	book.Chapters = chapters
	fmt.Printf("Book '%s' %d chapters, %d sections\n", bookName, len(chapters), book.SectionsCount())
	return book, nil
}
