---
Title: Variadic functions
Id: 1266
Score: 3
---
A variadic function can be called with any number of **trailing** arguments. Those elements are stored in a slice.

```go
package main

import "fmt"

func variadic(strs ...string) {
    // strs is a slice of string
    for i, str := range strs {
        fmt.Printf("%d: %s\n", i, str)
    }
}

func main() {
    variadic("Hello", "Goodbye")
    variadic("Str1", "Str2", "Str3")
}
```

[Playground](https://play.golang.org/p/rnzg1yK_Va)

You can also give a slice to a variadic function, with `...`:

```go
func main() {
    strs := []string {"Str1", "Str2", "Str3"}

    variadic(strs...)
}
```

[Playground](https://play.golang.org/p/gl4L5R9v8_)
