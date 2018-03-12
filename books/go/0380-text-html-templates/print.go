package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `print:   {{ print .Str .Num }}
println: {{ println .Str .Num }}
printf:  {{ printf "%s %#v %d" .Str .Str .Num }}
`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("print").Parse(tmplStr))

	data := struct {
		Str string
		Num int
	}{
		Str: "str",
		Num: 8,
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
