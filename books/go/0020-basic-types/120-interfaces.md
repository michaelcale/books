---
Title: Interfaces
Id: 24
SOId: 9010008c
---
An interface defines a set of methods on a struct.

Here's the definition of the [io.Reader](https://golang.org/pkg/io/#Reader) interface from the standard library:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

The smaller the interface, the better.

[Zero value](29) of interace is nil

Learn more about [interfaces](90).

