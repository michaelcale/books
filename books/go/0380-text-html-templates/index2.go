package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// :show start
	tmplStr := "Data: {{.}}\n"
	t := template.Must(template.New("simple").Parse(tmplStr))
	execWithData := func(data interface{}) {
		err := t.Execute(os.Stdout, data)
		if err != nil {
			log.Fatalf("t.Execute() failed with '%s'\n", err)
		}
	}

	execWithData(5)
	execWithData("foo")
	st := struct {
		Number int
		Str    string
	}{
		Number: 3,
		Str:    "hello",
	}
	execWithData(st)
	// :show end
}
