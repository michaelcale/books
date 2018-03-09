---
Title: Templates with custom logic
Id: 220
Score: 0
SOId: 13459
---
In this example, a function map named `funcMap` is supplied to the template via the `Funcs()` method and then invoked inside the template. Here, the function `increment()` is used to get around the lack of a less than or equal function in the templating language. Note in the output how the final item in the collection is handled.

A `-` at the beginning ``{{-`` or end ``-}}`` is used to trim whitespace and can be used to help make the template more legible.

```go
package main

import (
    "fmt"
    "os"
    "text/template"
)

var funcMap = template.FuncMap{
    "increment": increment,
}

func increment(x int) int {
    return x + 1
}

func main() {
    const (
        letter = `Dear {{with $names := .}}
        {{- range $i, $val := $names}}
            {{- if lt (increment $i) (len $names)}}
                {{- $val}}, {{else -}} and {{$val}}{{end}}
        {{- end}}{{end}}; How are you?`
    )

    tmpl, err := template.New("letter").Funcs(funcMap).Parse(letter)
    if err != nil {
        fmt.Println(err.Error())
    }

    tmpl.Execute(os.Stdout, []string{"Harry", "Jane", "Lisa", "George"})
}
```

Results in:

```text
Dear Harry, Jane, Lisa, and George; How are you?
```
