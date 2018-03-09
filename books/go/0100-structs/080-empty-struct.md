---
Title: Empty struct
Id: 88
SOId: 20546
---
A struct is a sequence of named elements, called fields, each of which has a name and a type. Empty struct has no fields, like this anonymous empty struct:

```go
var s struct{}
```

Or like this named empty struct type:

```go
type T struct{}
```

The interesting thing about the empty struct is that, its size is zero:

    fmt.Println(unsafe.Sizeof(s))

This prints `0`, so the empty struct itself takes no memory. so it is good option for quit channel, like:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan struct{})
    go func() {
        time.Sleep(1 * time.Second)
        close(done)
    }()

    fmt.Println("Wait...")
    <-done
    fmt.Println("done.")
}
```
