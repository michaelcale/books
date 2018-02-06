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

gBookTocSearchData = [
	["${chapter or aticle id"}, "${title}", "${synonim 1}", "${synanonim 2}", ...],
];

It's saved in wwww/essential/${bookname}/toc_search.js

TODO: optimize this as a one long string to search instead of multiple strings
in an array. Use 0 to separate chapter/article. Use 1 to separate title
from synonims.

Also, have original and lower-cased version of the string. We search lower-cased
but show the original. That avoids lowercasing during search.
*/

// TODO: add synonyms
func genBookTOCSearchMust(book *Book) {
	var toc [][]string
	for _, chapter := range book.Chapters {
		title := strings.TrimSpace(chapter.Title)
		a := []string{chapter.FileNameBase, title}
		toc = append(toc, a)
		for _, article := range chapter.Articles {
			title := strings.TrimSpace(article.Title)
			a := []string{article.FileNameBase, title}
			toc = append(toc, a)
		}
	}
	d, err := json.MarshalIndent(toc, "", "  ")
	u.PanicIfErr(err)
	s := "gBookTocSearchData = " + string(d) + ";"
	d = []byte(s)
	path := filepath.Join(destEssentialDir, book.FileNameBase, "toc_search.js")
	u.CreateDirForFile(path)
	err = ioutil.WriteFile(path, d, 0644)
	u.PanicIfErr(err)
	fmt.Printf("%s\n", path)
}
