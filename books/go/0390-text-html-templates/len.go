package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `len(nil)       : {{ len .SliceNil }}
len(emptySlice): {{ len .SliceEmpty }}
len(slice)     : {{ len .Slice }}
len(map)       : {{ len .Map }}
`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("len").Parse(tmplStr))

	data := struct {
		SliceNil   []int
		SliceEmpty []string
		Slice      []bool
		Map        map[int]bool
	}{
		SliceNil:   nil,
		SliceEmpty: []string{},
		Slice:      []bool{true, true, false},
		Map:        map[int]bool{5: true, 3: false},
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
