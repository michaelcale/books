Title: Single item template
Id: 13457
Score: 0
Body:
Note the use of `{{.}}` to output the item within the template.

    package main

    import (
        "fmt"
        "os"
        "text/template"
    )

    func main() {
        const (
            letter = `Dear {{.}}, How are you?`
        )

        tmpl, err := template.New("letter").Parse(letter)
        if err != nil {
            fmt.Println(err.Error())
        }

        tmpl.Execute(os.Stdout, "Professor Jones")
    }

Results in:

    Dear Professor Jones, How are you?

|======|
