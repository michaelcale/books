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
	default:
		log.Fatalf("unknown template '%s'\n", name)
	}
	return loadTemplateHelperMust(name, ref)
}

func execTemplateToFileMust(name string, data interface{}, path string) {
	fmt.Printf("%s\n", path)
	createDirForFileMust(path)
	tmpl := loadTemplateMust(name)
	f, err := os.Create(path)
	u.PanicIfErr(err)
	defer f.Close()
	err = tmpl.Execute(f, data)
	u.PanicIfErr(err)
}

func genIndex(books []*Book) {
	d := struct {
		Books []*Book
	}{
		Books: books,
	}
	path := filepath.Join("books_html", "index.html")
	execTemplateToFileMust("index.tmpl.html", d, path)
}

func genBookSection(section *Section) {
	path := section.destFilePath()
	execTemplateToFileMust("section.tmpl.html", section, path)
}

func genBookChapter(chapter *Chapter) {
	for _, section := range chapter.Sections {
		genBookSection(section)
	}
}

func genBook(book *Book) {
	// generate index.html for the book
	path := filepath.Join(book.DestDir, "index.html")
	execTemplateToFileMust("book_index.tmpl.html", book, path)
	for _, chapter := range book.Chapters {
		genBookChapter(chapter)
	}
	genBookTOCJSONMust(book)
}
