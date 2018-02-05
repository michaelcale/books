---
Title: Iota
Id: 2865
---
## Introduction
Iota provides a way of declaring numeric constants from a starting value that grows monotonically. Iota can be used to declare bitmasks which are often used in system and network programming and other lists of constants with related values.

## Remarks
The `iota` identifier is used to assign values to lists of constants. When iota is used in a list it starts with a value of zero, and increments by one for each value in the list of constants and is reset on each `const` keyword. Unlike the enumerations of other languages, iota can be used in expressions (eg. `iota + 1`) which allows for greater flexibility.
