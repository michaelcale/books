package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/essentialbooks/books/pkg/common"
	"github.com/kjk/u"
)

var (
	flgAnalytics      string
	flgPreview        bool
	allBookDirs       []string
	soUserIDToNameMap map[int]string
)

func parseFlags() {
	flag.StringVar(&flgAnalytics, "analytics", "", "google analytics code")
	flag.BoolVar(&flgPreview, "preview", false, "if true, will start watching for file changes and re-build everything")
	flag.Parse()
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

func copyCoversMust() {
	createDirForFileMaybeMust(filepath.Join("www", "covers", "foo.png"))
	fileInfos, err := ioutil.ReadDir("covers")
	u.PanicIfErr(err)
	for _, fi := range fileInfos {
		if fi.IsDir() || !fi.Mode().IsRegular() {
			continue
		}
		name := fi.Name()
		if strings.Contains(name, "@2x") {
			continue
		}
		dst := filepath.Join("www", "covers", name)
		if pathExists(dst) {
			continue
		}
		src := filepath.Join("covers", name)
		copyFile(dst, src)

	}
}

func genAllBooks() {
	timeStart := time.Now()

	copyCoversMust()

	var books []*Book
	for _, bookName := range allBookDirs {
		book, err := parseBook(bookName)
		maybePanicIfErr(err)
		if err != nil {
			continue
		}
		books = append(books, book)
	}

	copyToWwwMaybeMust(filepath.Join("tmpl", "main.css"))
	copyToWwwMaybeMust(filepath.Join("tmpl", "app.js"))
	genIndex(books)
	genIndexGrid(books)
	genAbout()

	nProcs := runtime.GOMAXPROCS(-1)
	sem := make(chan bool, nProcs)
	var wg sync.WaitGroup
	for _, book := range books {
		wg.Add(1)
		sem <- true
		go func(b *Book) {
			genBook(b)
			<-sem
			wg.Done()
			fmt.Printf("Generated %s, %d chapters, %d articles\n", b.Title, len(b.Chapters), b.ArticlesCount())
		}(book)
	}
	wg.Wait()
	fmt.Printf("Used %d procs, finished generating all book in %s\n", nProcs, time.Since(timeStart))
}

func loadSOUserMappingsMust() {
	path := filepath.Join("stack-overflow-docs-dump", "users.json.gz")
	err := common.JSONDecodeGzipped(path, &soUserIDToNameMap)
	u.PanicIfErr(err)
}

func main() {
	parseFlags()

	booksToImport := getBooksToImport(getBookDirs())
	for _, bookInfo := range booksToImport {
		allBookDirs = append(allBookDirs, bookInfo.NewName())
	}
	loadSOUserMappingsMust()
	os.RemoveAll("www")

	genAllBooks()
	if flgPreview {
		startPreview()
	}
}
