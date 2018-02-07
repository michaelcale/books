---
Title: Anonymous struct
Id: 1299
Score: 18
---
It is possible to create an anonymous struct:

```go
data := struct {
    Number int
    Text   string
} {
    42,
    "Hello world!",
}
```

Full example:

```go
package main

import (
    "fmt"
)

func main() {
    data := struct {Number int; Text string}{42, "Hello world!"} // anonymous struct
    fmt.Printf("%+v\n", data)
}
```

[Play it on playground](https://play.golang.org/p/atpNnP5wE_)
