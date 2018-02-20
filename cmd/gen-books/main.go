package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/essentialbooks/books/pkg/common"
	"github.com/kjk/u"
)

var (
	flgAnalytics          string
	flgPreview            bool
	flgUpdateGoPlayground bool
	allBookDirs           []string
	soUserIDToNameMap     map[int]string
)

func parseFlags() {
	flag.StringVar(&flgAnalytics, "analytics", "", "google analytics code")
	flag.BoolVar(&flgPreview, "preview", false, "if true will start watching for file changes and re-build everything")
	flag.BoolVar(&flgUpdateGoPlayground, "update-go-playground", false, "if true will upgrade links to go playground")
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

	copyToWwwStaticMaybeMust(filepath.Join("tmpl", "main.css"))
	copyToWwwStaticMaybeMust(filepath.Join("tmpl", "app.js"))
	copyToWwwStaticMaybeMust(filepath.Join("tmpl", "font-awesome.min.js"))
	genIndex(books)
	genIndexGrid(books)
	genAbout()

	for _, book := range books {
		genBook(book)
	}
	fmt.Printf("Used %d procs, finished generating all books in %s\n", getAlmostMaxProcs(), time.Since(timeStart))
}

func genAllBooks() {
	timeStart := time.Now()

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

	copyToWwwStaticMaybeMust(filepath.Join("tmpl", "main.css"))
	copyToWwwStaticMaybeMust(filepath.Join("tmpl", "app.js"))
	copyToWwwStaticMaybeMust(filepath.Join("tmpl", "font-awesome.min.js"))
	genIndex(books)
	genIndexGrid(books)
	genAbout()

	for _, book := range books {
		genBook(book)
	}
	fmt.Printf("Used %d procs, finished generating all books in %s\n", nProcs, time.Since(timeStart))
}

func loadSOUserMappingsMust() {
	path := filepath.Join("stack-overflow-docs-dump", "users.json.gz")
	err := common.JSONDecodeGzipped(path, &soUserIDToNameMap)
	u.PanicIfErr(err)
}

func timeFileCacheAndExit() {
	timeStart := time.Now()
	dir := filepath.Join("books", "go")
	err := cacheFilesInDir(dir)
	u.PanicIfErr(err)
	fmt.Printf("caching %d files in %s took %s\n", len(filePathToFileContent), dir, time.Since(timeStart))
	os.Exit(0)
}

func isBlacklistedForGetOutput(path string) bool {
	name := filepath.Base(path)
	switch name {
	case "timed_loop.go":
		return true
	}
	return false

}
func getOutputCb(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	if !strings.HasSuffix(path, ".go") {
		return nil
	}
	if isBlacklistedForGetOutput(path) {
		return nil
	}
	//fmt.Printf("%s\n", path)
	getOutput(path)
	return nil
}

func timeGetOutputAndExit() {
	timeStart := time.Now()
	dir := filepath.Join("books", "go")
	filepath.Walk(dir, getOutputCb)
	fmt.Printf("caching %d files in %s took %s\n", len(filePathToFileContent), dir, time.Since(timeStart))
	os.Exit(0)
}

func main() {
	parseFlags()

	if false {
		timeGetOutputAndExit()
	}

	if false {
		timeFileCacheAndExit()
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

	booksToImport := getBooksToImport(getBookDirs())
	for _, bookInfo := range booksToImport {
		allBookDirs = append(allBookDirs, bookInfo.NewName())
	}
	loadSOUserMappingsMust()
	os.RemoveAll("www")

	if flgPreview {
		go updateGoDeps()
	} else {
		updateGoDeps()
	}

	cacheFilesInDir("books")

	genAllBooks()
	if flgPreview {
		startPreview()
	}
}
