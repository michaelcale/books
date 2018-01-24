package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/kjk/u"
)

var (
	indexTmpl   *template.Template
	chapterTmpl *template.Template
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

// IndexTmplBook represents a book for index template
type IndexTmplBook struct {
	URL   string
	Title string
}

// IndexTmplModel represents data for index template
type IndexTmplModel struct {
	Books []*Book
}

func genIndex(books []*Book) {
	d := IndexTmplModel{
		Books: books,
	}
	path := filepath.Join("book_html", "index.html")
	execTemplateToFileMust("index.tmpl.html", d, path)
}
