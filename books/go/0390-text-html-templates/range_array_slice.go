package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `Elements of arrays or slice: {{ range . }}{{ . }} {{end}}
`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("range").Parse(tmplStr))

	array := [...]int{3, 8}
	err := t.Execute(os.Stdout, array)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}

	slice := []int{12, 5}
	err = t.Execute(os.Stdout, slice)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
