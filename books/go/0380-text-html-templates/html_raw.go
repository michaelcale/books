package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

// :show start
const tmplStr = `<div onlick="{{ .JS }}">{{ .HTML }}</div>
`

// :show end

func main() {
	// :show start

	html := template.Must(template.New("html").Parse(tmplStr))

	data := struct {
		JS   string
		HTML string
	}{
		JS:   `foo`,
		HTML: `<span>text</span>`,
	}

	fmt.Printf("Escaped:\n")
	err := html.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	fmt.Printf("\nUnescaped:\n")
	data2 := struct {
		JS   template.JS
		HTML template.HTML
	}{
		JS:   `foo`,
		HTML: `<span>text</span>`,
	}
	err = html.Execute(os.Stdout, data2)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	// :show end
}
