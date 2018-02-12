---
Title: Zero value of a map
Id: 2485
Score: 9
---
The zero value of a `map` is `nil` and has a length of `0`.

@file zero_value.go output

A `nil` map has no keys nor can keys be added. A `nil` map behaves like an empty map if read from but causes a runtime panic if written to.

@file zero_value2.go output allow_error

You should not read from or write to a zero value map.

