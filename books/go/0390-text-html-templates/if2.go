package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

// :show start
const tmplStr = `{{range . -}}
{{printf "%- 16s" .Name}} is: {{if .Value}}true{{else}}false{{end}}
{{end}}`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("if").Parse(tmplStr))

	var nilPtr *string = nil
	var nilSlice []float32
	emptySlice := []int{}

	data := []struct {
		Name  string
		Value interface{}
	}{
		{"bool false", false},
		{"bool true", true},
		{"integer 0", 0},
		{"integer 1", 1},
		{"float32 0", float32(0)},
		{"float64 NaN", math.NaN},
		{"empty string", ""},
		{"non-empty string", "haha"},
		{"nil slice", nilSlice},
		{"empty slice", emptySlice},
		{"non-empty slice", []int{3}},
		{"nil pointer", nilPtr},
	}

	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
