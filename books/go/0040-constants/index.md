---
Title: Constants
Id: 1047
---
Go supports constants of character, string, boolean, and numeric values.

Constant basics:
```go
// Greeting is an exported (public) string constant
const Greeting string = "Hello World"

// we can group const declarations
const (
    // years is an unexported (package private) int constant
    years int = 10
    truth bool = true
)
```

Constants can be used like variable, except for the fact that the value cannot be changed. Here's an example:
```go
package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s) // constant

    // A `const` statement can appear anywhere a `var` statement can.
    const n = 10
    fmt.Println(n)                           // 10
    fmt.Printf("n=%d is of type %T\n", n, n) // n=10 is of type int

    const m float64 = 4.3
    fmt.Println(m) // 4.3

    // An untyped constant takes the type needed by its context.
    // For example, here `math.Sin` expects a `float64`.
    const x = 10
    fmt.Println(math.Sin(x)) // -0.5440211108893699
}
```

[Playground](https://play.golang.org/p/MI48yM88dE)