package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `Slice[0]: {{ index .Slice 0 }}
SliceNested[1][0]: {{ index .SliceNested 1 0 }}
Map["key"]: {{ index .Map "key" }}
`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("index").Parse(tmplStr))

	data := struct {
		Slice       []string
		SliceNested [][]int
		Map         map[string]int
	}{
		Slice: []string{"first", "second"},
		SliceNested: [][]int{
			{3, 1},
			{2, 3},
		},
		Map: map[string]int{
			"key": 5,
		},
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
