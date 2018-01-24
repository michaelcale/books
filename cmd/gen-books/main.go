package main

import (
	"fmt"
	"time"
)

var bookDirs = []string{
	".NET Framework",
	"algorithm",
	"Android",
	"Angular 2",
	"AngularJS",
	"Bash",
	"C Language",
	"C++",
	"C# Language",
	"CSS",

	"jQuery",
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
	for _, book := range books {
		genBook(book)
		fmt.Printf("Generated %s, %d chapters, %d sections\n", book.Title, len(book.Chapters), book.SectionsCount())
	}

	fmt.Printf("Finished in %s\n", time.Since(timeStart))
}
