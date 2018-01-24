package main

import (
	"fmt"
	"time"
)

var bookDirs = []string{
	"jQuery",
}

func main() {
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
	}
}
