---
Title: go run
Id: 324
SOId: 17005
---
`go run` will run a program without creating an executable file. Mostly useful for development. `run` will only execute packages whose *package name* is **main**.

To demonstrate, we will use a simple Hello World example `main.go`:

```go
package main

import fmt

func main() {
    fmt.Println("Hello, World!")
}
```

Execute without compiling to a file: `go run main.go`

## Run multiple files in package

If program is split into multiple files, you must provide all files:

```sh
$ go run main.go assets.go
```

<!-- TODO: note about not being able to just run *.go when *_test.go files exists -->