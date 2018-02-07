---
Title: Conditional loop
Id: 3889
Score: 6
---
The `for` keyword is also used for conditional loops, traditionally `while` loops in other programming languages.

```go
package main

import (
    "fmt"
)

func main() {
    i := 0
    for i < 3 { // Will repeat if condition is true
        i++
        fmt.Println(i)
    }
}
```

[Playground](https://play.golang.org/p/18NqQo3PA6)

Will output:

```text
1
2
3
```

**infinite loop:**

```go
for {
    // This will run until a return or break.
}
```
