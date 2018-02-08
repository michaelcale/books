---
Title: Structures
Search: struct
Id: rd600098
---
A structure groups multiple values into a single entity.

Struct basics:
```go
type MyStruct struct {
    IntVal int
    StringVal string
    unexportedIntVal int
}
```

[Zero value](a-6069) of `struct` is a struct whose fields are set to thier respective zero values.

Learn more about [structs](ch-374).
