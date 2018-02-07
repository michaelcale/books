---
Title: Declaring a constant
Id: 3376
Score: 3
---
Constants are declared like variables, but using the `const` keyword:

```go
const Greeting string = "Hello World"
const Years int = 10
const Truth bool = true
```

Like for variables, names starting with an upper case letter are exported (_public_), names starting with lower case are not.

```go
// not exported
const alpha string = "Alpha"
// exported
const Beta string = "Beta"
```

Constants can be used like any other variable, except for the fact that the value cannot be changed. Here's an example:

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
