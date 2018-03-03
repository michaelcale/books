---
Title: Arrays
Id: 390
---

Arrays in Go have a fixed sized. They can't grow.

Because of that arrays in Go are rarely used. Instead [slices](a-733) are used in most cases.

[Zero value](a-6069) of array is an array where all values have zero value.

Elements of arrays are laid out in memory consecutively, which is good for speed.

Arrays are passed by value which means that passing an array argument to a function copies the whole array. This is slow if the array is large.

Array basics:

@file index.go output sha1:8ae8a8a7dc7cf2756dc2744232f4140f9bb7d633 goplayground:7Vg96smORkS
