---
Title: Arrays
Id: 390
---

Arrays in Go have fixed sized. They can't grow.

Because of that arrays in Go are used rarely. Instead [slices](ch-733) are used in most cases.

[Zero value](a-6069) of array is array where all values have zero value.

Elements of arrays are laid out in memory consequitevely, which is good for speed.

Arrays are passed by value which means that passing array argument to a function copies the whole array. This is slow if the array is large.

Array basics:

@file index.go output
