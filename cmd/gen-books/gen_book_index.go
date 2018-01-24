package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/kjk/u"
)

var (
	indexTmpl     *template.Template
	chapterTmpl   *template.Template
	bookIndexTmpl *template.Template
)

func tmplPath(name string) string {
	return filepath.Join("book_html", name)
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
	case "chapter.tmpl.html":
		ref = &chapterTmpl
	case "book_index.tmpl.html":
		ref = &bookIndexTmpl
	default:
		log.Fatalf("unknown template '%s'\n", name)
	}
	return loadTemplateHelperMust(name, ref)
}

func execTemplateToFileMust(name string, data interface{}, path string) {
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
	path := filepath.Join("book_html", "index.html")
	execTemplateToFileMust("index.tmpl.html", d, path)
}

func genBook(book *Book) {
	err := os.MkdirAll(book.DestDir, 0755)
	u.PanicIfErr(err)

	// generate index.html for the book
	path := filepath.Join(book.DestDir, "index.html")
	execTemplateToFileMust("book_index.tmpl.html", book, path)

}
