package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `js escape  : {{ js .JS }}
html escape: {{ html .HTML }}
url escape : {{ urlquery .URL }}
`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("print").Parse(tmplStr))

	data := struct {
		JS   string
		HTML string
		URL  string
	}{
		JS:   `function me(s) { return "foo"; }`,
		HTML: `<div>text</div>`,
		URL:  `http://www.programming-books.io`,
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
