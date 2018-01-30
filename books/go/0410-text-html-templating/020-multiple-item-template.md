Title: Multiple item template
Id: 13458
Score: 0
Body:
Note the use of `{{range .}}` and `{{end}}` to cycle over the collection.

    package main

    import (
        "fmt"
        "os"
        "text/template"
    )

    func main() {
        const (
            letter = `Dear {{range .}}{{.}}, {{end}} How are you?`
        )

        tmpl, err := template.New("letter").Parse(letter)
        if err != nil {
            fmt.Println(err.Error())
        }

        tmpl.Execute(os.Stdout, []string{"Harry", "Jane", "Lisa", "George"})
    }

Results in:

    Dear Harry, Jane, Lisa, George,  How are you?
|======|
