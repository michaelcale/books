---
Title: Hello, World!
Id: 833
---
Create file `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

[Playground](https://play.golang.org/p/I3l_5RKJts)

Compile and run with `go run hello.go`.

## Output:

```text
Hello, 世界
```

Once you are happy with the code it can be compiled to an executable with `go build hello.go`.

On Windows this will create `hello.exe` executable.

On every other OS, this will create `./hello` executable.
