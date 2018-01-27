package main

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/kjk/programming-books/pkg/common"
)

func getBooksToImport() []*common.Book {
	var res []*common.Book
	for _, bookInfo := range common.BooksToProcess {
		if !bookInfo.Import {
			continue
		}
		res = append(res, bookInfo)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].NewName() < res[j].NewName()
	})
	return res
}

// TODO: probably more
func getDefaultLangForBook(bookName string) string {
	s := strings.ToLower(bookName)
	switch s {
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

func main() {
	timeStart := time.Now()
	var books []*Book
	booksToImport := getBooksToImport()
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
			fmt.Printf("Generated %s, %d chapters, %d sections\n", b.Title, len(b.Chapters), b.SectionsCount())
		}(book)
	}
	wg.Wait()

	fmt.Printf("Used %d procs, finished in %s\n", nProcs, time.Since(timeStart))
}
