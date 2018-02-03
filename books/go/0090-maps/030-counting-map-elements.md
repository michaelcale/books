---
Title: Counting map elements
Id: 2528
Score: 14
---
The built-in function [`len`](https://golang.org/pkg/builtin/#len) returns the number of elements in a `map`

    m := map[string]int{}
    len(m) // 0

    m["foo"] = 1
    len(m) // 1

If a variable points to a `nil` map, then `len` returns 0.

    var m map[string]int
    len(m) // 0
