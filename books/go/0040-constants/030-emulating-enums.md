---
Title: Emulating enums
Id: 801000c4
---
Go doesn't have a syntax for enumerations, but you can emulate it with constants and a naming scheme.

Consider a C++ enum:
```c++
enum {
    tagHtml,
    tagBody,
    taDiv
};
```

In Go you can do it as:
```go
const (
    tagBody = iota,
    tagDiv
    tagHr
)
```

## Adding type safety

In the above example `tagBody` etc. is an untyped constant so it can be assigned to any number type.

We can define a unique type for our enum:
```go
type HTMLTag int

const (
    tagBody HTMLTag = iota,
    tagDiv
    tagHr
)
```
