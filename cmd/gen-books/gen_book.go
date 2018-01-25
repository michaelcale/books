package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/kjk/u"
)

var (
	indexTmpl     *template.Template
	bookIndexTmpl *template.Template
	chapterTmpl   *template.Template
	sectionTmpl   *template.Template
	aboutTmpl     *template.Template

	gitHubBaseURL = "https://github.com/kjk/programming-books"
)

func createDirForFileMust(path string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	u.PanicIfErr(err)
}

func tmplPath(name string) string {
	return filepath.Join("books_html", name)
}

func loadTemplateHelperMust(name string, ref **template.Template) *template.Template {
	res := *ref
	if res != nil {
		return res
	}
	path := tmplPath(name)
	t, err := template.ParseFiles(path)
	u.PanicIfErr(err)
	*ref = t
	return t
}

func loadTemplateMust(name string) *template.Template {
	var ref **template.Template
	switch name {
	case "index.tmpl.html":
		ref = &indexTmpl
	case "book_index.tmpl.html":
		ref = &bookIndexTmpl
	case "chapter.tmpl.html":
		ref = &chapterTmpl
	case "section.tmpl.html":
		ref = &sectionTmpl
	case "about.tmpl.html":
		ref = &aboutTmpl
	default:
		log.Fatalf("unknown template '%s'\n", name)
	}
	return loadTemplateHelperMust(name, ref)
}

func execTemplateToFileSilentMust(name string, data interface{}, path string) {
	createDirForFileMust(path)
	tmpl := loadTemplateMust(name)
	f, err := os.Create(path)
	u.PanicIfErr(err)
	defer f.Close()
	err = tmpl.Execute(f, data)
	u.PanicIfErr(err)
}

func execTemplateToFileMust(name string, data interface{}, path string) {
	fmt.Printf("%s\n", path)
	execTemplateToFileSilentMust(name, data, path)
}

func genIndex(books []*Book) {
	d := struct {
		Books      []*Book
		GitHubText string
		GitHubURL  string
	}{
		Books:      books,
		GitHubText: "GitHub",
		GitHubURL:  gitHubBaseURL,
	}
	path := filepath.Join("books_html", "index.html")
	execTemplateToFileMust("index.tmpl.html", d, path)
}

func genAbout() {
	path := filepath.Join("books_html", "about.html")
	execTemplateToFileMust("about.tmpl.html", nil, path)
}

func genBookSection(section *Section) {
	if section.BodyHTML == "" {
		html := markdownToHTML([]byte(section.BodyMarkdown))
		section.BodyHTML = template.HTML(html)
	}
	path := section.destFilePath()
	execTemplateToFileSilentMust("section.tmpl.html", section, path)
}

func genBookChapter(chapter *Chapter) {
	for _, section := range chapter.Sections {
		genBookSection(section)
	}

	path := chapter.destFilePath()
	execTemplateToFileSilentMust("chapter.tmpl.html", chapter, path)
}

func setCurrentChapter(chapters []*Chapter, current int) {
	for i, chapter := range chapters {
		chapter.IsCurrent = current == i
	}
}

func genBook(book *Book) {
	// generate index.html for the book
	path := filepath.Join(book.DestDir, "index.html")
	execTemplateToFileSilentMust("book_index.tmpl.html", book, path)
	for i, chapter := range book.Chapters {
		setCurrentChapter(book.Chapters, i)
		genBookChapter(chapter)
	}
	//genBookTOCJSONMust(book)
}
