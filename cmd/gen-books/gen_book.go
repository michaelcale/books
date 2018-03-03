package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	// top-level directory where .html files are generated
	destDir = "www"
	tmplDir = "tmpl"
)

var ( // directory where generated .html files for books are
	destEssentialDir  = filepath.Join(destDir, "essential")
	pathAppJS         = "/s/app.js"
	pathFontAwesomeJS = "/s/font-awesome.min.js"
	pathMainCSS       = "/s/main.css"
)

var (
	templateNames = []string{
		"index.tmpl.html",
		"index-grid.tmpl.html",
		"book_index.tmpl.html",
		"chapter.tmpl.html",
		"article.tmpl.html",
		"about.tmpl.html",
		"feedback.tmpl.html",
	}
	templates = make([]*template.Template, len(templateNames))

	gitHubBaseURL = "https://github.com/essentialbooks/books"
	siteBaseURL   = "https://www.programming-books.io"
)

func unloadTemplates() {
	templates = make([]*template.Template, len(templateNames))
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
	for i, tmplName := range templateNames {
		if tmplName == name {
			ref = &templates[i]
			break
		}
	}
	if ref == nil {
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
		Books             []*Book
		GitHubText        string
		GitHubURL         string
		Analytics         template.HTML
		PathAppJS         string
		PathFontAwesomeJS string
		PathMainCSS       string
	}{
		Books:             books,
		GitHubText:        "GitHub",
		GitHubURL:         gitHubBaseURL,
		Analytics:         googleAnalytics,
		PathAppJS:         pathAppJS,
		PathFontAwesomeJS: pathFontAwesomeJS,
		PathMainCSS:       pathMainCSS,
	}
	path := filepath.Join(destDir, "index.html")
	execTemplateToFileMaybeMust("index.tmpl.html", d, path)
}

func genIndexGrid(books []*Book) {
	d := struct {
		Books             []*Book
		Analytics         template.HTML
		PathAppJS         string
		PathFontAwesomeJS string
		PathMainCSS       string
	}{
		Books:             books,
		Analytics:         googleAnalytics,
		PathAppJS:         pathAppJS,
		PathFontAwesomeJS: pathFontAwesomeJS,
		PathMainCSS:       pathMainCSS,
	}
	path := filepath.Join(destDir, "index-grid.html")
	execTemplateToFileMaybeMust("index-grid.tmpl.html", d, path)
}

func genFeedback() {
	d := struct {
		Analytics         template.HTML
		PathAppJS         string
		PathFontAwesomeJS string
		PathMainCSS       string
	}{
		Analytics:         googleAnalytics,
		PathAppJS:         pathAppJS,
		PathFontAwesomeJS: pathFontAwesomeJS,
		PathMainCSS:       pathMainCSS,
	}
	fmt.Printf("writing feedback.html\n")
	path := filepath.Join(destDir, "feedback.html")
	execTemplateToFileMaybeMust("feedback.tmpl.html", d, path)
}

func genAbout() {
	d := struct {
		Analytics         template.HTML
		PathAppJS         string
		PathFontAwesomeJS string
		PathMainCSS       string
	}{
		Analytics:         googleAnalytics,
		PathAppJS:         pathAppJS,
		PathFontAwesomeJS: pathFontAwesomeJS,
		PathMainCSS:       pathMainCSS,
	}
	fmt.Printf("writing about.html\n")
	path := filepath.Join(destDir, "about.html")
	execTemplateToFileMaybeMust("about.tmpl.html", d, path)
}

func genArticle(article *Article, currChapNo int) {
	addSitemapURL(article.CanonnicalURL())

	d := struct {
		*Article
		CurrentChapterNo  int
		Analytics         template.HTML
		PathAppJS         string
		PathFontAwesomeJS string
		PathMainCSS       string
	}{
		Article:           article,
		CurrentChapterNo:  currChapNo,
		Analytics:         googleAnalytics,
		PathAppJS:         pathAppJS,
		PathFontAwesomeJS: pathFontAwesomeJS,
		PathMainCSS:       pathMainCSS,
	}

	path := article.destFilePath()
	execTemplateToFileSilentMaybeMust("article.tmpl.html", d, path)
}

func genChapter(chapter *Chapter, currNo int) {
	addSitemapURL(chapter.CanonnicalURL())
	for _, article := range chapter.Articles {
		genArticle(article, currNo)
	}

	path := chapter.destFilePath()
	d := struct {
		*Chapter
		CurrentChapterNo  int
		Analytics         template.HTML
		PathAppJS         string
		PathFontAwesomeJS string
		PathMainCSS       string
	}{
		Chapter:           chapter,
		CurrentChapterNo:  currNo,
		Analytics:         googleAnalytics,
		PathAppJS:         pathAppJS,
		PathFontAwesomeJS: pathFontAwesomeJS,
		PathMainCSS:       pathMainCSS,
	}
	execTemplateToFileSilentMaybeMust("chapter.tmpl.html", d, path)

	for _, imagePath := range chapter.images {
		imageName := filepath.Base(imagePath)
		dst := chapter.destImagePath(imageName)
		copyFileMaybeMust(dst, imagePath)
	}
}

func genBook(book *Book) {
	fmt.Printf("Started genering book %s\n", book.Title)
	timeStart := time.Now()

	// generate index.html for the book
	err := os.MkdirAll(book.destDir, 0755)
	maybePanicIfErr(err)
	if err != nil {
		return
	}

	path := filepath.Join(book.destDir, "index.html")
	d := struct {
		Book              *Book
		Analytics         template.HTML
		PathAppJS         string
		PathFontAwesomeJS string
		PathMainCSS       string
	}{
		Book:              book,
		Analytics:         googleAnalytics,
		PathAppJS:         pathAppJS,
		PathFontAwesomeJS: pathFontAwesomeJS,
		PathMainCSS:       pathMainCSS,
	}

	execTemplateToFileSilentMaybeMust("book_index.tmpl.html", d, path)

	addSitemapURL(book.CanonnicalURL())

	for i, chapter := range book.Chapters {
		book.sem <- true
		book.wg.Add(1)
		go func(idx int, chap *Chapter) {
			genChapter(chap, idx)
			book.wg.Done()
			<-book.sem
		}(i+1, chapter)
	}
	genBookTOCSearchMust(book)
	book.wg.Wait()

	fmt.Printf("Generated %s, %d chapters, %d articles in %s\n", book.Title, len(book.Chapters), book.ArticlesCount(), time.Since(timeStart))
}
