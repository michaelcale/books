---
Title: Characters and runes
Search: char, rune
Id: 9010002e
---
Go has 2 types of characters:
* `byte` is a 1 byte value, an alias for `uint8` type
* `rune` is a 4 byte Unicode code-point, an alias for `int32` type

[Zero value](a-6069) of a `byte` and `rune` is 0.

## Iterate over string using bytes

@file characters.go output sha1:3da66fb5895a3e079aae5bc307daa026bc933ce1 goplayground:9JMcrAOALyg

## Iterate over string using runes

@file characters-2.go output sha1:86731d891f071d94c0ad19802423f4265473df70 goplayground:qLPCaibSrZC

Note: When iterating over runes, the index will be the offset within the string where the rune begins, rather than the index of the rune itself.
