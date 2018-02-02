Title: Hello, World!
Id: 833
Score: 92
Body:
Create file `hello.go` with content:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}
```

[Playground](https://play.golang.org/p/I3l_5RKJts)

When `Go` is [installed correctly](a-20381) this program can be compiled and run with `go run hello.go`.

# Output:

```text
Hello, 世界
```

Once you are happy with the code it can be compiled to an executable with `go build hello.go`.

On Windows this will create `hello.exe` executable.

On every other OS, this will create `./hello` executable.
