package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `Elements of map:
{{ range $k, $v := . }}{{ $k }}: {{ $v }}
{{end}}`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("range").Parse(tmplStr))

	m := map[string]int{
		"one":  1,
		"five": 5,
	}
	err := t.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
