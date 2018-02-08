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
```go
a := [3]int{4, 5}     // array of 2 ints

// access element of array
fmt.Printf("a[2]: %d\n", a[2])

// set element of array
a[1] = 3

// get size of array
fmt.Printf("size of array a: %d\n", len(a))

// when using [...] size will be deduced from { ... }
a2 := [...]{4, 8, -1} // array of 3 integers
```
