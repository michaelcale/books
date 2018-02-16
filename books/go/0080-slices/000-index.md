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

[Zero value of a slice](a-6807) is nil.

Basic of slices:

@file index.go output sha1:0aad3a082470e1c324f760dbf3b616843908d2b2 goplayground:nuqqi6Y0fF6
