package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
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
	index2Tmpl    *template.Template
	bookIndexTmpl *template.Template
	chapterTmpl   *template.Template
	articleTmpl   *template.Template
	aboutTmpl     *template.Template

	gitHubBaseURL = "https://github.com/essentialbooks/books"
)

func unloadTemplates() {
	indexTmpl = nil
	bookIndexTmpl = nil
	chapterTmpl = nil
	articleTmpl = nil
	aboutTmpl = nil
}

func tmplPath(name string) string {
	return filepath.Join(tmplDir, name)
}

func loadTemplateHelperMaybeMust(name string, ref **template.Template) *template.Template {
	res := *ref
	if res != nil {
		return res
	}
	path := tmplPath(name)
	//fmt.Printf("loadTemplateHelperMust: %s\n", path)
	t, err := template.ParseFiles(path)
	maybePanicIfErr(err)
	if err != nil {
		return nil
	}
	*ref = t
	return t
}

func loadTemplateMaybeMust(name string) *template.Template {
	var ref **template.Template
	switch name {
	case "index.tmpl.html":
		ref = &indexTmpl
	case "index2.tmpl.html":
		ref = &index2Tmpl
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
	return loadTemplateHelperMaybeMust(name, ref)
}

func execTemplateToFileSilentMaybeMust(name string, data interface{}, path string) {
	tmpl := loadTemplateMaybeMust(name)
	if tmpl == nil {
		return
	}
	f, err := os.Create(path)
	maybePanicIfErr(err)
	defer f.Close()
	err = tmpl.Execute(f, data)
	maybePanicIfErr(err)
}

func execTemplateToFileMaybeMust(name string, data interface{}, path string) {
	execTemplateToFileSilentMaybeMust(name, data, path)
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
	execTemplateToFileMaybeMust("index.tmpl.html", d, path)
}

// Cover represents a book cover
type Cover struct {
	URL string
}

func genIndex2() {
	var covers []Cover
	for _, coverName := range coverNames {
		uri := "/covers/" + coverName + ".png"
		c := Cover{
			URL: uri,
		}
		covers = append(covers, c)
	}
	d := struct {
		Covers        []Cover
		AnalyticsCode string
	}{
		Covers:        covers,
		AnalyticsCode: flgAnalytics,
	}
	path := filepath.Join(destDir, "index2.html")
	execTemplateToFileMaybeMust("index2.tmpl.html", d, path)

}

func genAbout() {
	d := struct {
		AnalyticsCode string
	}{
		AnalyticsCode: flgAnalytics,
	}
	path := filepath.Join(destDir, "about.html")
	execTemplateToFileMaybeMust("about.tmpl.html", d, path)
}

func genBookArticle(article *Article) {
	article.AnalyticsCode = flgAnalytics
	// TODO: move as a method on Article
	if article.BodyHTML == "" {
		defLang := getDefaultLangForBook(article.Book().Title)
		html := markdownToHTML([]byte(article.BodyMarkdown), defLang, article.Book())
		article.BodyHTML = template.HTML(html)
	}
	path := article.destFilePath()
	execTemplateToFileSilentMaybeMust("article.tmpl.html", article, path)
}

func genBookChapter(chapter *Chapter) {
	for _, article := range chapter.Articles {
		genBookArticle(article)
	}

	path := chapter.destFilePath()
	chapter.AnalyticsCode = flgAnalytics
	execTemplateToFileSilentMaybeMust("chapter.tmpl.html", chapter, path)
}

func setCurrentChapter(chapters []*Chapter, current int) {
	for i, chapter := range chapters {
		chapter.IsCurrent = current == i
	}
}

func genBook(book *Book) {
	// generate index.html for the book
	err := os.MkdirAll(book.destDir, 0755)
	maybePanicIfErr(err)
	if err != nil {
		return
	}

	path := filepath.Join(book.destDir, "index.html")
	book.AnalyticsCode = flgAnalytics
	execTemplateToFileSilentMaybeMust("book_index.tmpl.html", book, path)
	for i, chapter := range book.Chapters {
		setCurrentChapter(book.Chapters, i)
		genBookChapter(chapter)
	}
	//genBookTOCJSONMust(book)
}
