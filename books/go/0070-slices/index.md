---
Title: Slices
Id: 733
---

Slices in Go are used where most languages would use arrays as they are a growable sequence of values of the same type.

Memory used by slice is provided by an [fixed size array](ch-390). A slice is a view into that array.

Slice has a length and capacity.

Capacity represents how many total elements a slice can have. That's the size of underlying array.

Length is the current number of elements in the slice.

The difference between capacity and length is how many elements we can append to a slice before we have to re-allocate underlying array.

[Zero value](a-6069) of slice is nil.


## Syntax
 - slice := make([]type, len, cap) // create a new slice
 - slice = append(slice, item) // append a item to a slice
 - slice = append(slice, items...) // append slice of items to a slice
 - len := len(slice) // get the length of a slice
 - cap := cap(slice) // get the capacity of a slice
 - elNum := copy(dst, slice) // copy a the contents of a slice to an other slice
