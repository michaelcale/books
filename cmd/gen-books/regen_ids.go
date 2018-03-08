package main

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/essentialbooks/books/pkg/kvstore"
	"github.com/kjk/u"
)

// this is a one-time operation to regenerate ids to make them shorter
func regenIDSAndExit() {
	booksToImport := getBooksToImport(getBookDirs())
	for _, bookInfo := range booksToImport {
		allBookDirs = append(allBookDirs, bookInfo.NewName())
	}
	loadSOUserMappingsMust()
	currID := 0
	for _, bookName := range allBookDirs {
		book, err := parseBook(bookName)
		u.PanicIfErr(err)
		for _, chapter := range book.Chapters {
			if chapter.FileNameBase == "contributors" {
				continue
			}
			currID++

			path := chapter.indexFilePath
			doc := chapter.indexDoc
			saveDocWithNewID(path, doc, currID)

			for _, article := range chapter.Articles {
				currID++

				path := article.sourceFilePath
				doc, err := kvstore.ParseKVFile(path)
				u.PanicIfErr(err)
				saveDocWithNewID(path, doc, currID)
			}
		}
	}
	os.Exit(0)
}

func saveDocWithNewID(path string, doc kvstore.Doc, newID int) error {
	id, err := doc.Get("Id")
	u.PanicIfErr(err)
	doc = kvstore.ReplaceOrAppend(doc, "SOId", id)
	id = strconv.Itoa(newID)
	doc = kvstore.ReplaceOrAppend(doc, "Id", id)
	return saveDoc(path, doc)
}

func saveDoc(path string, doc kvstore.Doc) error {
	s, err := kvstore.SerializeDoc(doc)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(s), 0644)
}
