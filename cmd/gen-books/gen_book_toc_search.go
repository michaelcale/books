package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/kjk/u"
)

/*
Generates a javascript file that looks like:

gBookToc = [
	[${chapter or aticle id}, ${parentIdx}, ${title}, ${synonim 1}, ${synonim 2}, ...],
];

It's saved in wwww/essential/${bookname}/toc_search.js

TODO: optimize this as a one long string to search instead of multiple strings
in an array. Use 0 to separate chapter/article. Use 1 to separate title
from synonims.

Also, have original and lower-cased version of the string. We search lower-cased
but show the original. That avoids lowercasing during search.
*/

func genBookTOCSearchMust(book *Book) {
	var toc [][]interface{}
	for _, chapter := range book.Chapters {
		title := strings.TrimSpace(chapter.Title)
		tocItem := []interface{}{chapter.FileNameBase, -1, title}
		toc = append(toc, tocItem)
		chapIdx := len(toc) - 1
		u.PanicIf(chapIdx < 0)

		headings := chapter.Headings()
		for _, heading := range headings {
			tocItem = []interface{}{"", chapIdx, heading}
			toc = append(toc, tocItem)
		}

		for _, article := range chapter.Articles {
			title := strings.TrimSpace(article.Title)
			tocItem = []interface{}{article.FileNameBase, chapIdx, title}
			for _, syn := range article.SearchSynonyms {
				tocItem = append(tocItem, syn)
			}
			toc = append(toc, tocItem)

			headings := article.Headings()
			articleIdx := len(toc) - 1
			for _, heading := range headings {
				tocItem = []interface{}{"", articleIdx, heading}
				toc = append(toc, tocItem)
			}
		}
	}
	d, err := json.MarshalIndent(toc, "", "  ")
	u.PanicIfErr(err)
	s := "gBookToc = " + string(d) + ";"
	d = []byte(s)
	path := filepath.Join(destEssentialDir, book.FileNameBase, "toc_search.js")
	u.CreateDirForFile(path)
	err = ioutil.WriteFile(path, d, 0644)
	u.PanicIfErr(err)
	fmt.Printf("%s\n", path)
}
