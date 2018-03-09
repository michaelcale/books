---
Title: Multiple item template
Id: 219
Score: 0
SOId: 13458
---
Note the use of `{{range .}}` and `{{end}}` to cycle over the collection.

```go
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
```

Results in:

```text
Dear Harry, Jane, Lisa, George,  How are you?
```
