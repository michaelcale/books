package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
var tmplStr = `Data from a field: '{{ .Field }}'
Data from a method: '{{ .Method }}'
`

// :show end

type Data struct {
	Field int
}

func (d Data) Method() string {
	return "data from a method"
}

func main() {
	// :show start
	t := template.New("method")
	t, err := t.Parse(tmplStr)
	if err != nil {
		log.Fatalf("template.Parse() failed with '%s'\n", err)
	}

	data := Data{
		Field: 5,
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
