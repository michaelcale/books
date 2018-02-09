---
Title: Arrays
Search: array
Id: rd60004c
---
Arrays in Go have fixed sized. They can't grow.

They are used rarely. Instead in most cases we use [slices](a-rd6000rd).

A slice is growable and implemented as a view into its underlying array.

Array basics:

@file arrays.go output

[Zero value](a-6069) of array is array where all values have zero value.

Learn more about [arrays](ch-390).
