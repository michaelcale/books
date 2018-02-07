---
Title: Slices are Pointers to Array Segments
Id: 23177
Score: 0
---
Slices are **pointers** to arrays, with the length of the segment, and its capacity. They behave as pointers, and assigning their value to another slice, will assign the memory address. To **copy** a slice value to another, use the built-in **copy** function: `func copy(dst, src []Type) int` (returns the amount of items copied).

```go
package main

import (
    "fmt"
)

func main() {
    x := []byte{'a', 'b', 'c'}
    fmt.Printf("%s", x)       // prints: abc
    y := x
    y[0], y[1], y[2] = 'x', 'y', 'z'
    fmt.Printf("%s", x)       // prints: xyz
    z := make([]byte, len(x))
    // To copy the value to another slice, but
    // but not the memory address use copy:
    _ = copy(z, x)            // returns count of items copied
    fmt.Printf("%s", z)       // prints: xyz
    z[0], z[1], z[2] = 'a', 'b', 'c'
    fmt.Printf("%s", x)       // prints: xyz
    fmt.Printf("%s", z)       // prints: abc
}
```
