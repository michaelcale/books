---
Title: Panic
Id: 17009
Score: 1
---
A panic halts normal execution flow and exits the current function.  Any deferred calls will then be executed before control is passed to the next higher function on the stack.  Each stack's function will exit and run deferred calls until the panic is handled using a deferred `recover()`, or until the panic reaches `main()` and terminates the program.  If this occurs, the argument provided to panic and a stack trace will be printed to `stderr`.

```go
package main

import "fmt"

func foo() {
    defer fmt.Println("Exiting foo")
    panic("bar")
}

func main() {
    defer fmt.Println("Exiting main")
    foo()
}
```

Output:

```text
Exiting foo
Exiting main
panic: bar


goroutine 1 [running]:
panic(0x128360, 0x1040a130)
    /usr/local/go/src/runtime/panic.go:481 +0x700
main.foo()
    /tmp/sandbox550159908/main.go:7 +0x160
main.main()
    /tmp/sandbox550159908/main.go:12 +0x120
```

It is important to note that `panic` will accept any type as its parameter.
