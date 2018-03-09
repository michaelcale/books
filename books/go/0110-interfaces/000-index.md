---
Title: Interfaces
Id: 90
SOId: 1221
---
An interface describes a set of methods on a type.

It's similar to `interface` type in C#.

Interfaces are used to abstract behavior.

For example a standard library defines `io.Reader` interface:

```go
type Reader interface {
    Read(d []byte) (int, error)
}
```

Most functions that operate on streams of binary data (e.g. json decoder) take `io.Reader` as a source of data. That way we can implement `Reader` interface for physical files, bytes in memory, network connections and have `json.Decode` work on all those sources.

Here's how to define and implement a simple interface:

@file index.go output sha1:4211ae2b94a999a48fd954197938d1b8dec96086 goplayground:ba83kFxk4Yg

Unlike most other languages, interfaces are satisfied implicitly.

We don't have to explicitly declare that `struct User` implements interface `Stringer`.

Interfaces can only contain methods, not data. You can use [struct embedding](84) if you want to re-use both methods and data.

You can only define methods on types defined in the same package. We had to define type alias `MyInt` because we can't add methods to built-int type `int`.

<!-- TODO: how interfaces are implemented internally -->

<!-- TODO: is MyInt type alias or something else? -->

<!-- TODO: implementation of interfaces -->