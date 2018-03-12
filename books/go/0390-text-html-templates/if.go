package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `{{range . -}}
{{if .IsNew}}'{{.Name}}' is new{{else}}'{{.Name}}' is not new{{end}}
{{end}}`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("if").Parse(tmplStr))

	data := []struct {
		Name  string
		IsNew bool
	}{
		{"Bridge", false},
		{"Electric battery", true},
	}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
