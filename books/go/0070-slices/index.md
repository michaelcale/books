---
Title: Slices
Id: 733
---

## Introduction
A slice is a data structure that encapsulates an array so that the programmer can add as many elements as needed without having to worry about memory management. Slices can be cut into sub-slices very efficiently, since the resulting slices all point to the same internal array. Go programmers often take advantage of this to avoid copying arrays, which would typically be done in many other programming languages.

## Syntax
 - slice := make([]type, len, cap) // create a new slice
 - slice = append(slice, item) // append a item to a slice
 - slice = append(slice, items...) // append slice of items to a slice
 - len := len(slice) // get the length of a slice
 - cap := cap(slice) // get the capacity of a slice
 - elNum := copy(dst, slice) // copy a the contents of a slice to an other slice
