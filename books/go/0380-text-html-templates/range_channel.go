package main

import (
	"log"
	"os"
	"text/template"
)

// :show start
const tmplStr = `Elements of a channel: {{ range . }}{{ . }} {{end}}
`

// :show end

func main() {
	// :show start
	t := template.Must(template.New("range").Parse(tmplStr))

	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	err := t.Execute(os.Stdout, ch)
	if err != nil {
		log.Fatalf("t.Execute() failed with '%s'\n", err)
	}
	// :show end
}
