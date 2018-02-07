---
Title: Defining separate behaviours in different platforms
Id: 26889
Score: 0
---
Different platforms can have separate implementations of the same method. This example also illustrates how build tags and file suffixes can be used together.

File `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World from Conditional Compilation Doc!")
    printDetails()
}
```

`details.go`:

```go
// +build !windows

package main

import "fmt"

func printDetails() {
    fmt.Println("Some specific details that cannot be found on Windows")
}
```

`details_windows.go`:
```go
package main

import "fmt"

func printDetails() {
    fmt.Println("Windows specific details")
}
```
