---
Title: Characters and runes
Search: char, rune
Id: 9010002e
---
Go has 2 types of characters:
* `byte` is 1 byte value, an alias for `uint8` type
* `rune` is 4 byte Unicode code-point, an alias for `int32` type

[Zero value](a-6069) of a `byte` and `rune` is 0.

## Iterate over string using bytes

@file characters.go output

## Iterate over string using runes

@file characters-2.go output

Note: When iterating as runes, you get an index within strings where rune starts, not a rune number within the string.