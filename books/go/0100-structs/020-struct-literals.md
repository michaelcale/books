---
Title: Struct literals
Id: 82
SOId: 12466
---
A value of a struct type can be written using a *struct literal* that specifies values for its fields.

```go
type Point struct { X, Y int }
p := Point{1, 2}
```

The above example specifies every field in the right order. Which is not useful, because programmers have to remember the exact fields in order. More often, a struct can be initialized by listing some or all of the field names and their corresponding values.

```go
    anim := gif.GIF{LoopCount: nframes}
```

Omitted fields are set to the zero value for its type.

Note: **The two forms cannot be mixed in the same literal.**
