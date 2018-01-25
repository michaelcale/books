package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var bookDirs = []string{
	".NET Framework",
	//"algorithm",
	//"Android",
	//"Angular 2",
	//"AngularJS",
	//"Bash",
	"C Language",
	"C++",
	"C# Language",
	"CSS",
	"jQuery",
	"Go",
}

func main() {
	timeStart := time.Now()
	var books []*Book
	for _, bookName := range bookDirs {
		timeStart := time.Now()
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
