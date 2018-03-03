---
Title: Arrays
Search: array
Id: rd60004c
---
Arrays in Go have a fixed sized. They can't grow.

They are used rarely. Instead in most cases we use [slices](a-rd6000rd).

A slice is growable and implemented as a view into its underlying array.

Array basics:

@file arrays.go output sha1:8824b8dc01bf1d44d35c30053c7a4fb2d049c257 goplayground:sg72vy6F-ci

[Zero value](a-6069) of array is array where all values have zero value.

Learn more about [arrays](a-390).
