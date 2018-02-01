package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/essentialbooks/books/pkg/common"
	"github.com/kjk/u"
)

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

func main() {
	timeStart := time.Now()
	var books []*Book
	booksToImport := getBooksToImport(getBookDirs())
	for _, b := range booksToImport {
		fmt.Printf("Importing book: %s\n", b.NewName())
	}

	os.RemoveAll("www")

	for _, bookInfo := range booksToImport {
		timeStart := time.Now()
		bookName := bookInfo.NewName()
		book, err := parseBook(bookName)
		if err != nil {
			fmt.Printf("Error '%s' parsing book '%s'\n", err, bookName)
			return
		}
		books = append(books, book)
		fmt.Printf("Generating book '%s' took %s\n", bookName, time.Since(timeStart))
	}
	genIndex(books)
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

	fmt.Printf("Used %d procs, finished in %s\n", nProcs, time.Since(timeStart))
}
