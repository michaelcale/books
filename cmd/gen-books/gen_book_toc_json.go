package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/kjk/u"
)

// BookArticleTOC represents an article in a book TOC
type BookArticleTOC struct {
	Title string `json:"title"`
	// TODO: url?
}

// BookChapterTOC represents chapter in book TOC
type BookChapterTOC struct {
	Title    string           `json:"title"`
	Articles []BookArticleTOC `json:"article"`
}

// BookTOC represents table of contents of a book
type BookTOC struct {
	Name     string           `json:"name"`
	Chapters []BookChapterTOC `json:"chapters"`
}

func genBookTOCJSONData(book *Book) ([]byte, error) {
	bookTOC := BookTOC{
		Name: book.Title,
	}
	for _, ch := range book.Chapters {
		var articles []BookArticleTOC
		for _, a := range ch.Articles {
			sectoc := BookArticleTOC{
				Title: a.Title,
			}
			articles = append(articles, sectoc)
		}
		chtoc := BookChapterTOC{
			Title:    ch.Title,
			Articles: articles,
		}
		bookTOC.Chapters = append(bookTOC.Chapters, chtoc)
	}
	return json.MarshalIndent(&bookTOC, "", "  ")
}

func genBookTOCJSONMust(book *Book) {
	d, err := genBookTOCJSONData(book)
	u.PanicIfErr(err)
	path := filepath.Join("books_html", "book", book.FileNameBase, "toc.json")
	u.CreateDirForFile(path)
	err = ioutil.WriteFile(path, d, 0644)
	u.PanicIfErr(err)
	fmt.Printf("%s\n", path)
}
