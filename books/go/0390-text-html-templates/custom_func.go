package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `5 + 5 = {{ sum 5 .Arg }}
`

// :show end

func sum(x, y int) int {
	return x + y
}

func main() {
	// :show start
	customFunctions := template.FuncMap{
		"sum": sum,
	}

	t := template.Must(template.New("func").Funcs(customFunctions).Parse(tmplStr))

	data := struct {
		Arg int
	}{
		Arg: 5,
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
