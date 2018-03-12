package main

import (
	"log"
	"os"
	"text/template"
)

type Data struct {
	True  bool
	False bool
}

// :show start
const tmplStr = `Or:  {{ if or .True .False }}true{{ else }}false{{ end }}
And: {{ if and .True .False }}true{{ else }}false{{ end }}
Not: {{ if not .False }}true{{ else }}false{{ end }}
`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("and_or_not").Parse(tmplStr))

	data := Data{True: true, False: false}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
