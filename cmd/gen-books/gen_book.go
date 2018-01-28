package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/kjk/u"
)

const (
	// top-level directory where .html files are generated
	destDir = "www"
	tmplDir = "tmpl"
)

var ( // directory where generated .html files for books are
	destEssentialDir = filepath.Join(destDir, "essential")
)

var (
	indexTmpl     *template.Template
	bookIndexTmpl *template.Template
	chapterTmpl   *template.Template
	articleTmpl   *template.Template
	aboutTmpl     *template.Template

	gitHubBaseURL = "https://github.com/kjk/programming-books"
)

func createDirForFileMust(path string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0755)
	u.PanicIfErr(err)
}

func tmplPath(name string) string {
	return filepath.Join(tmplDir, name)
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
	case "article.tmpl.html":
		ref = &articleTmpl
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
	path := filepath.Join(destDir, "index.html")
	execTemplateToFileMust("index.tmpl.html", d, path)
}

func genAbout() {
	path := filepath.Join(destDir, "about.html")
	execTemplateToFileMust("about.tmpl.html", nil, path)
}

func genBookArticle(article *Article) {
	// TODO: move as a method on Article
	if article.BodyHTML == "" {
		defLang := getDefaultLangForBook(article.Book().Title)
		html := markdownToHTML([]byte(article.BodyMarkdown), defLang)
		article.BodyHTML = template.HTML(html)
	}
	path := article.destFilePath()
	execTemplateToFileSilentMust("article.tmpl.html", article, path)
}

func genBookChapter(chapter *Chapter) {
	for _, article := range chapter.Articles {
		genBookArticle(article)
	}

	path := chapter.destFilePath()
	execTemplateToFileSilentMust("chapter.tmpl.html", chapter, path)
}

func setCurrentChapter(chapters []*Chapter, current int) {
	for i, chapter := range chapters {
		chapter.IsCurrent = current == i
	}
}

func copyFileMust(dst, src string) {
	createDirForFileMust(dst)

	in, err := os.Open(src)
	u.PanicIfErr(err)
	defer in.Close()
	out, err := os.Create(dst)
	u.PanicIfErr(err)
	defer out.Close()
	_, err = io.Copy(out, in)
	u.PanicIfErr(err)
}

func copyCSSMust() {
	src := filepath.Join(tmplDir, "main.css")
	dst := filepath.Join(destDir, "main.css")
	copyFileMust(dst, src)
}

func genBook(book *Book) {
	// generate index.html for the book
	copyCSSMust()
	path := filepath.Join(book.destDir, "index.html")
	execTemplateToFileSilentMust("book_index.tmpl.html", book, path)
	for i, chapter := range book.Chapters {
		setCurrentChapter(book.Chapters, i)
		genBookChapter(chapter)
	}
	//genBookTOCJSONMust(book)
}
