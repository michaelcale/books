package main

import (
	"fmt"
	html_template "html/template"
	"log"
	"os"
	text_template "text/template"
)

// :show start
const tmplStr = `<div onlick="{{ .JS }}">{{ .HTML }}</div>
`

// :show end

func main() {
	// :show start
	txt := text_template.Must(text_template.New("text").Parse(tmplStr))

	html := html_template.Must(html_template.New("html").Parse(tmplStr))

	data := struct {
		JS   string
		HTML string
		URL  string
	}{
		JS:   `foo`,
		HTML: `<span>text</span>`,
		URL:  `http://www.programming-books.io`,
	}

	err := txt.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	fmt.Println()

	err = html.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	// :show end
}
