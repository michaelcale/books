---
Title: Zero value of a map
Id: 2485
---
The zero value of a `map` is `nil` and has a length of `0`.

@file zero_value.go output sha1:2fdb095b28d8b58ad0bde153735ecec4c4201ea4 goplayground:MfIwOym0yKy

A `nil` map has no keys nor can keys be added. A `nil` map behaves like an empty map if read from, but causes a runtime panic if written to.

@file zero_value2.go output allow_error sha1:e57d417eae9ce61d41605f144137d25a3fe3c0b0 goplayground:ALhSLNlHunv

You should not read from or write to a zero value map.

