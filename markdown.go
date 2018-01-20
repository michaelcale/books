package main

import (
	"github.com/kjk/programming-books/pkg/mdutil"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func mdToHTML(d []byte) []byte {
	//r := blackfriday.NewHTMLRenderer()
	return blackfriday.Run(d)
}

func mdFmt(src []byte, defaultLang string) ([]byte, error) {
	opts := &mdutil.Options{DefaultLang: defaultLang}
	return mdutil.Process(src, opts)
}
