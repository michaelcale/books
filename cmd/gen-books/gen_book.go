package main

import (
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

	gitHubBaseURL = "https://github.com/essentialbooks/books"
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
	//fmt.Printf("loadTemplateHelperMust: %s\n", path)
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
	tmpl := loadTemplateMust(name)
	f, err := os.Create(path)
	u.PanicIfErr(err)
	defer f.Close()
	err = tmpl.Execute(f, data)
	u.PanicIfErr(err)
}

func execTemplateToFileMust(name string, data interface{}, path string) {
	execTemplateToFileSilentMust(name, data, path)
}

func genIndex(books []*Book) {
	d := struct {
		Books         []*Book
		GitHubText    string
		GitHubURL     string
		AnalyticsCode string
	}{
		Books:         books,
		GitHubText:    "GitHub",
		GitHubURL:     gitHubBaseURL,
		AnalyticsCode: flgAnalytics,
	}
	path := filepath.Join(destDir, "index.html")
	execTemplateToFileMust("index.tmpl.html", d, path)
}

func genAbout() {
	d := struct {
		AnalyticsCode string
	}{
		AnalyticsCode: flgAnalytics,
	}
	path := filepath.Join(destDir, "about.html")
	execTemplateToFileMust("about.tmpl.html", d, path)
}

func genBookArticle(article *Article) {
	article.AnalyticsCode = flgAnalytics
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
	chapter.AnalyticsCode = flgAnalytics
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
	err := os.MkdirAll(book.destDir, 0755)
	u.PanicIfErr(err)

	path := filepath.Join(book.destDir, "index.html")
	book.AnalyticsCode = flgAnalytics
	execTemplateToFileSilentMust("book_index.tmpl.html", book, path)
	for i, chapter := range book.Chapters {
		setCurrentChapter(book.Chapters, i)
		genBookChapter(chapter)
	}
	//genBookTOCJSONMust(book)
}
