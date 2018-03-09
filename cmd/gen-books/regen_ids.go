package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/essentialbooks/books/pkg/kvstore"
	"github.com/kjk/u"
)

type modifiedDoc struct {
	doc  kvstore.Doc
	path string
}

// this is a one-time operation to regenerate ids to make them shorter
func regenIDSAndExit() {
	// maps old StackOverflow id to new ids
	idMap := make(map[string]string)
	var docs []*modifiedDoc
	booksToImport := getBooksToImport(getBookDirs())
	for _, bookInfo := range booksToImport {
		allBookDirs = append(allBookDirs, bookInfo.NewName())
	}
	loadSOUserMappingsMust()
	for _, bookName := range allBookDirs {
		book, err := parseBook(bookName)
		u.PanicIfErr(err)
		for _, chapter := range book.Chapters {
			if chapter.FileNameBase == "contributors" {
				continue
			}

			path := chapter.indexFilePath
			doc, err := kvstore.ParseKVFile(path)
			u.PanicIfErr(err)

			mdoc := &modifiedDoc{
				doc:  doc,
				path: path,
			}
			docs = append(docs, mdoc)

			for _, article := range chapter.Articles {
				path := article.sourceFilePath
				doc, err := kvstore.ParseKVFile(path)
				u.PanicIfErr(err)

				mdoc := &modifiedDoc{
					doc:  doc,
					path: path,
				}
				docs = append(docs, mdoc)
			}
		}
	}

	// assign new ids
	for id, mdoc := range docs {
		doc := mdoc.doc
		soID, err := doc.Get("Id")
		u.PanicIfErr(err)
		doc = kvstore.ReplaceOrAppend(doc, "SOId", soID)
		newID := strconv.Itoa(id + 1)
		doc = kvstore.ReplaceOrAppend(doc, "Id", newID)
		idMap[soID] = newID
	}

	for _, mdoc := range docs {
		doc = mdoc.doc
		body := doc.GetSilent("Body", "")
		if body == "" {
			continue
		}
		body = fixLinks(body, idMap)
		doc = kvstore.ReplaceOrAppend(doc, "Body", body)
		err = saveDoc(mdoc.path, doc)
		u.PanicIfErr(err)
	}
	os.Exit(0)
}

// fix links (${oldID}) => (${newID})
func fixLinks(s string, idMap map[string]string) string {
	for k, v := range idMap {
		old := "(" + k + ")"
		new := "(" + v + ")"
		s = strings.Replace(s, old, new, -1)
	}
	return s
}

func saveDoc(path string, doc kvstore.Doc) error {
	s, err := kvstore.SerializeDoc(doc)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(s), 0644)
}
