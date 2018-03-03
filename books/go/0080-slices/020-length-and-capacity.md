---
Title: Length and capacity
Id: 3561
---
Slices have both length and capacity.  The length of a slice is the number of elements *currently* in the slice, while the capacity is the number of elements the slice *can hold* before needing to be reallocated.

When creating a slice using the built-in `make()` function, you can specify its length, and optionally its capacity. If the capacity is not explicitly specified, it will default to the value of the specified length.
```
var s = make([]int, 3, 5) // length 3, capacity 5
```

You can check the length of a slice with the built-in `len()` function:
```
var n = len(s) // n == 3
```

You can check the capacity with the built-in `cap()` function:
```
var c = cap(s) // c == 5
```

Elements created by `make()` are set to the zero value for the element type of the slice:
```
for idx, val := range s {
    fmt.Println(idx, val)
}
// output:
// 0 0
// 1 0
// 2 0
```


You cannot access elements beyond the length of a slice, even if the index is within capacity:
```
var x = s[3] // panic: runtime error: index out of range
```
However, as long as the capacity exceeds the length, you can append new elements without reallocating:
```
var t = []int{3, 4}
s = append(s, t) // s is now []int{0, 0, 0, 3, 4}
n = len(s) // n == 5
c = cap(s) // c == 5
```

If you append to a slice which lacks the capacity to accept the new elements, the underlying array will be reallocated for you with sufficient capacity:
```
var u = []int{5, 6}
s = append(s, u) // s is now []int{0, 0, 0, 3, 4, 5, 6}
n = len(s) // n == 7
c = cap(s) // c > 5
```
It is, therefore, generally good practice to allocate sufficient capacity when first creating a slice, if you know how much space you'll need, to avoid unnecessary reallocations.
