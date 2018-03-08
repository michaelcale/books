package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
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
	destEssentialDir       = filepath.Join(destDir, "essential")
	pathAppJS              = "/s/app.js"
	pathMainCSS            = "/s/main.css"
	pathFaviconICO         = "/s/favicon.ico"
	totalHTMLBytes         int
	totalHTMLBytesMinified int
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
		"404.tmpl.html",
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
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	maybePanicIfErr(err)

	d := buf.Bytes()
	if doMinify {
		d2, err := minifier.Bytes("text/html", d)
		maybePanicIfErr(err)
		if err == nil {
			totalHTMLBytes += len(d)
			totalHTMLBytesMinified += len(d2)
			d = d2
		}
	}
	err = ioutil.WriteFile(path, d, 0644)
	maybePanicIfErr(err)
}

func execTemplateToFileMaybeMust(name string, data interface{}, path string) {
	execTemplateToFileSilentMaybeMust(name, data, path)
}

// PageCommon is a common information for most pages
type PageCommon struct {
	Analytics      template.HTML
	PathAppJS      string
	PathMainCSS    string
	PathFaviconICO string
}

func getPageCommon() PageCommon {
	return PageCommon{
		Analytics:      googleAnalytics,
		PathAppJS:      pathAppJS,
		PathMainCSS:    pathMainCSS,
		PathFaviconICO: pathFaviconICO,
	}
}

func gen404(books []*Book) {
	d := struct {
		PageCommon
		GitHubText string
		GitHubURL  string
	}{
		PageCommon: getPageCommon(),
		GitHubText: "GitHub",
		GitHubURL:  gitHubBaseURL,
	}
	path := filepath.Join(destDir, "404.html")
	execTemplateToFileMaybeMust("404.tmpl.html", d, path)
}

func genIndex(books []*Book) {
	d := struct {
		PageCommon
		Books      []*Book
		GitHubText string
		GitHubURL  string
	}{
		PageCommon: getPageCommon(),
		Books:      books,
		GitHubText: "GitHub",
		GitHubURL:  gitHubBaseURL,
	}
	path := filepath.Join(destDir, "index.html")
	execTemplateToFileMaybeMust("index.tmpl.html", d, path)
}

func genIndexGrid(books []*Book) {
	d := struct {
		PageCommon
		Books []*Book
	}{
		PageCommon: getPageCommon(),
		Books:      books,
	}
	path := filepath.Join(destDir, "index-grid.html")
	execTemplateToFileMaybeMust("index-grid.tmpl.html", d, path)
}

func genFeedback() {
	d := getPageCommon()
	fmt.Printf("writing feedback.html\n")
	path := filepath.Join(destDir, "feedback.html")
	execTemplateToFileMaybeMust("feedback.tmpl.html", d, path)
}

func genAbout() {
	d := getPageCommon()
	fmt.Printf("writing about.html\n")
	path := filepath.Join(destDir, "about.html")
	execTemplateToFileMaybeMust("about.tmpl.html", d, path)
}

func genArticle(article *Article, currChapNo int) {
	addSitemapURL(article.CanonnicalURL())

	d := struct {
		PageCommon
		*Article
		CurrentChapterNo int
	}{
		PageCommon:       getPageCommon(),
		Article:          article,
		CurrentChapterNo: currChapNo,
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
		PageCommon
		*Chapter
		CurrentChapterNo int
	}{
		PageCommon:       getPageCommon(),
		Chapter:          chapter,
		CurrentChapterNo: currNo,
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

	genBookTOCSearchMust(book)

	// generate index.html for the book
	err := os.MkdirAll(book.destDir, 0755)
	maybePanicIfErr(err)
	if err != nil {
		return
	}

	path := filepath.Join(book.destDir, "index.html")
	d := struct {
		PageCommon
		Book *Book
	}{
		PageCommon: getPageCommon(),
		Book:       book,
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
	book.wg.Wait()

	fmt.Printf("Generated %s, %d chapters, %d articles in %s\n", book.Title, len(book.Chapters), book.ArticlesCount(), time.Since(timeStart))
}
