---
Title: go build
Id: 325
SOId: 17006
---
`go build` will compile a program into an executable file.

To demonstrate, we will use a simple Hello World example main.go:

```go
package main

import fmt

func main() {
    fmt.Println("Hello, World!")
}
```

To compile the program run: `go build main.go`

`build` creates an executable program, in this case: `main` or `main.exe`. You can then run this file to see the output `Hello, World!`. You can also copy it to a similar system that doesn't have Go installed, *make it executable*, and run it there.

## Specify OS or Architecture in build:

You can specify what system or architecture to build by modifying the `env` before `build`:

```sh
env GOOS=linux go build main.go # builds for Linux
env GOARCH=arm go build main.go # builds for ARM architecture
```

## Build multiple files

If your package is split into multiple files **and** the package name is **main** (that is, *it is not an importable package*), you must specify all the files to build:

    go build main.go assets.go # outputs an executable: main

## Building a package

To build a package called `main`, you can simply use:

    go build . # outputs an executable with name as the name of enclosing folder
