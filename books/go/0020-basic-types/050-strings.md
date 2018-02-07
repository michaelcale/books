---
Title: Strings
Search: string
Id: rd6000k1
---
Unlike languages like Python, C#, Java or Swift, a string in Go is simply an immutable sequence of bytes (8-bit `byte` values).

[Zero value](a-6069) of a `string` type is an empty string.

Basic string usage:
```go
var s string // empty string ""
s1 := "string\nliteral\nwith\tescape characters"
s2 := `raw string literal
which doesnt't recgonize escape characters like \n
`
fmt.Printf("sum of string: %s\n", s + s1 + s2)
```

Learn more about [strings](ch-9666).
