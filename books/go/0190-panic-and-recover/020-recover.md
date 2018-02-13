---
Title: Recover
Id: 17010
Score: 1
---
Recover as the name implies, can attempt to recover from a `panic`.
The recover *must* be attempted in a deferred statement as normal execution flow has been halted.  The `recover` statement must appear *directly* within the deferred function enclosure.  Recover statements in functions called by deferred function calls will not be honored.  The `recover()` call will return the argument provided to the initial panic, if the program is currently panicking.  If the program is not currently panicking, `recover()` will return `nil`.

```go
package main

import "fmt"

func foo() {
    panic("bar")
}

func bar() {
    defer func() {
        if msg := recover(); msg != nil {
            fmt.Printf("Recovered with message %s\n", msg)
        }
    }()
    foo()
    fmt.Println("Never gets executed")
}

func main() {
    fmt.Println("Entering main")
    bar()
    fmt.Println("Exiting main the normal way")
}
```

Output:

```text
Entering main
Recovered with message bar
Exiting main the normal way
```
