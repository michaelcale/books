Title: Templates with structs
Id: 13460
Score: 0
Body:
Note how field values are obtained using `{{.FieldName}}`.

    package main

    import (
        "fmt"
        "os"
        "text/template"
    )

    type Person struct {
        FirstName string
        LastName  string
        Street    string
        City      string
        State     string
        Zip       string
    }

    func main() {
        const (
            letter = `------------------------------
    {{range .}}{{.FirstName}} {{.LastName}}
    {{.Street}}
    {{.City}}, {{.State}} {{.Zip}}

    Dear {{.FirstName}},
        How are you?

    ------------------------------
    {{end}}`
        )

        tmpl, err := template.New("letter").Parse(letter)
        if err != nil {
            fmt.Println(err.Error())
        }

        harry := Person{
            FirstName: "Harry",
            LastName:  "Jones",
            Street:    "1234 Main St.",
            City:      "Springfield",
            State:     "IL",
            Zip:       "12345-6789",
        }

        jane := Person{
            FirstName: "Jane",
            LastName:  "Sherman",
            Street:    "8511 1st Ave.",
            City:      "Dayton",
            State:     "OH",
            Zip:       "18515-6261",
        }

        tmpl.Execute(os.Stdout, []Person{harry, jane})
    }

Results in:

    ------------------------------
    Harry Jones
    1234 Main St.
    Springfield, IL 12345-6789
    
    Dear Harry,
        How are you?
    
    ------------------------------
    Jane Sherman
    8511 1st Ave.
    Dayton, OH 18515-6261
    
    Dear Jane,
        How are you?
    
    ------------------------------
|======|
