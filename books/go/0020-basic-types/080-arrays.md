---
Title: Arrays
Search: array
Id: rd60004c
---
Arrays in Go have fixed sized. They can't grow.

Array basics:
```go
var a1 = [2]byte{3, 8}  // array of 2 bytes
// when using [...] size will be deduced from { ... }
a2 := [...]int{1, 2, 3} // array of 3 integers
```

Arrays in Go are used rarely.

A [slice](a-rd6000rd) is a view into an array that can grow its underlying array.

For that reason slices are used as frequently as arrays in languages like Python or Java.

[Zero value](a-6069) of array is array where all values have zero value.

Learn more about [arrays](ch-390).
