package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/essentialbooks/books/pkg/common"
	"github.com/kjk/u"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

var (
	flgAnalytics          string
	flgPreview            bool
	flgUpdateGoPlayground bool
	allBookDirs           []string
	soUserIDToNameMap     map[int]string
	googleAnalytics       template.HTML
	doMinify              bool
	minifier              *minify.M
)

const (
	// https://www.netlify.com/docs/headers-and-basic-auth/#custom-headers
	netlifyHeaders = `
# long-lived caching
/s/*
  Cache-Control: max-age=31536000
/*
  X-Content-Type-Options: nosniff
  X-Frame-Options: DENY
  X-XSS-Protection: 1; mode=block
`
)

const (
	googleAnalyticsTmpl = `<script async src="https://www.googletagmanager.com/gtag/js?id=%s"></script>
    <script>
        window.dataLayer = window.dataLayer || [];
        function gtag(){dataLayer.push(arguments);}
        gtag('js', new Date());
        gtag('config', '%s');
    </script>
`
)

func parseFlags() {
	flag.StringVar(&flgAnalytics, "analytics", "", "google analytics code")
	flag.BoolVar(&flgPreview, "preview", false, "if true will start watching for file changes and re-build everything")
	flag.BoolVar(&flgUpdateGoPlayground, "update-go-playground", false, "if true will upgrade links to go playground")
	flag.Parse()

	if flgAnalytics != "" {
		s := fmt.Sprintf(googleAnalyticsTmpl, flgAnalytics, flgAnalytics)
		googleAnalytics = template.HTML(s)
	}
}

func dirFromBook(book *common.Book) string {
	return common.MakeURLSafe(book.NewName())
}

func isBookImported(bookDirs []string, book *common.Book) bool {
	dir := dirFromBook(book)
	for _, dir2 := range bookDirs {
		if dir == dir2 {
			return true
		}
	}
	return false
}

func getBooksToImport(bookDirs []string) []*common.Book {
	var res []*common.Book
	for _, book := range common.BooksToProcess {
		if isBookImported(bookDirs, book) {
			res = append(res, book)
		}
	}
	return res
}

// TODO: probably more
func getDefaultLangForBook(bookName string) string {
	s := strings.ToLower(bookName)
	switch s {
	case "go":
		return "go"
	case "android":
		return "java"
	case "ios":
		return "ObjectiveC"
	case "microsoft sql server":
		return "sql"
	case "node.js":
		return "javascript"
	case "mysql":
		return "sql"
	case ".net framework":
		return "c#"
	}
	return s
}

func getBookDirs() []string {
	dirs, err := common.GetDirs("books")
	u.PanicIfErr(err)
	return dirs
}

func shouldCopyImage(path string) bool {
	return !strings.Contains(path, "@2x")
}

func copyFilesRecur(dstDir, srcDir string, shouldCopyFunc func(path string) bool) {
	createDirMust(dstDir)
	fileInfos, err := ioutil.ReadDir(srcDir)
	u.PanicIfErr(err)
	for _, fi := range fileInfos {
		name := fi.Name()
		if fi.IsDir() {
			dst := filepath.Join(dstDir, name)
			src := filepath.Join(srcDir, name)
			copyFilesRecur(dst, src, shouldCopyFunc)
			continue
		}

		src := filepath.Join(srcDir, name)
		dst := filepath.Join(dstDir, name)
		shouldCopy := true
		if shouldCopyFunc != nil {
			shouldCopy = shouldCopyFunc(src)
		}
		if !shouldCopy {
			continue
		}
		if pathExists(dst) {
			continue
		}
		copyFileMust(dst, src)
	}
}

func copyCoversMust() {
	copyFilesRecur(filepath.Join("www", "covers"), "covers", shouldCopyImage)
}

func getAlmostMaxProcs() int {
	// leave some juice for other programs
	nProcs := runtime.NumCPU() - 2
	if nProcs < 1 {
		return 1
	}
	return nProcs
}

func genSelectedBooks(bookDirs []string) {
	fmt.Printf("genSelectedBooks: %+v\n", bookDirs)
	timeStart := time.Now()

	var books []*Book
	for _, bookName := range bookDirs {
		book, err := parseBook(bookName)
		maybePanicIfErr(err)
		if err != nil {
			continue
		}
		book.sem = make(chan bool, getAlmostMaxProcs())
		books = append(books, book)
	}
	fmt.Printf("Parsed books in %s\n", time.Since(timeStart))

	copyToWwwStaticMaybeMust("main.css")
	copyToWwwStaticMaybeMust("app.js")
	genIndex(books)
	genIndexGrid(books)
	genAbout()
	genFeedback()

	for _, book := range books {
		genBook(book)
	}
	fmt.Printf("Used %d procs, finished generating all books in %s\n", getAlmostMaxProcs(), time.Since(timeStart))
}

func genAllBooks() {
	timeStart := time.Now()
	clearSitemapURLS()
	copyCoversMust()

	nProcs := getAlmostMaxProcs()

	var books []*Book
	for _, bookName := range allBookDirs {
		book, err := parseBook(bookName)
		maybePanicIfErr(err)
		if err != nil {
			continue
		}
		book.sem = make(chan bool, nProcs)
		books = append(books, book)
	}
	fmt.Printf("Parsed books in %s\n", time.Since(timeStart))

	copyToWwwStaticMaybeMust("main.css")
	copyToWwwStaticMaybeMust("app.js")
	genIndex(books)
	genIndexGrid(books)
	genAbout()
	genFeedback()

	for _, book := range books {
		genBook(book)
	}
	writeSitemap()
	fmt.Printf("Used %d procs, finished generating all books in %s\n", nProcs, time.Since(timeStart))
}

func loadSOUserMappingsMust() {
	path := filepath.Join("stack-overflow-docs-dump", "users.json.gz")
	err := common.JSONDecodeGzipped(path, &soUserIDToNameMap)
	u.PanicIfErr(err)
}

func main() {

	parseFlags()

	if false {
		updateGoDeps()
		//updateGoDepsInFile("books/go/0250-time-date/strftime.go")
		os.Exit(0)
	}

	if false {
		genTwitterImagesAndExit()
	}

	if false {
		testGetGoPlaygroundShareIDAndExit()
	}

	if flgUpdateGoPlayground {
		goBookDir := filepath.Join("books", "go")
		updateGoPlaygroundLinks(goBookDir)
		os.Exit(0)
	}

	minifier = minify.New()
	minifier.AddFunc("text/css", css.Minify)
	minifier.AddFunc("text/html", html.Minify)
	minifier.AddFunc("text/javascript", js.Minify)
	// less aggresive minification because html validators
	// report this as html errors
	minifier.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
		KeepEndTags:      true,
	})

	booksToImport := getBooksToImport(getBookDirs())
	for _, bookInfo := range booksToImport {
		allBookDirs = append(allBookDirs, bookInfo.NewName())
	}
	loadSOUserMappingsMust()

	os.RemoveAll("www")
	createDirMust(filepath.Join("www", "s"))
	path := filepath.Join("www", "_headers")
	err := ioutil.WriteFile(path, []byte(netlifyHeaders), 0644)
	u.PanicIfErr(err)

	updateGoDeps()
	cacheFilesInDir("books")

	doMinify = !flgPreview

	clearErrors()
	genAllBooks()
	printAndClearErrors()
	if flgPreview {
		startPreview()
	}
}
