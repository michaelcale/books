---
Title: Command line arguments
Id: 14046
Score: 1
---
Command line arguments parsing is Go is very similar to other languages. In you code you just access slice of arguments where first argument will be the name of program itself.

Quick example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {

    progName := os.Args[0]
    arguments := os.Args[1:]

    fmt.Printf("Here we have program '%s' launched with following flags: ", progName)

    for _, arg := range arguments {
        fmt.Printf("%s ", arg)
    }

    fmt.Println("")
}
```

And output would be:

```sh
$ ./cmd test_arg1 test_arg2
Here we have program './cmd' launched with following flags: test_arg1 test_arg2
```

Each argument is just a string. In `os` package it looks like:
`var Args []string`
