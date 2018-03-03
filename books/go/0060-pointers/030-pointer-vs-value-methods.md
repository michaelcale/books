---
Title: Pointer vs. value methods
Id: 6049
Score: 1
---

## Pointer Methods

Pointer methods can be called even if the variable is itself not a pointer.

According to the [Go Spec](https://golang.org/ref/spec#Method_values),

>  . . . a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`.

You can see this in this example:

@file pointer_methods.go output sha1:db09cfa9a5d39dd638eda6b7eab0e658c0642dee goplayground:hHqYV-X605Z

## Value Methods

Similarly to pointer methods, value methods can be called even if the variable is itself not a value.

According to the [Go Spec](https://golang.org/ref/spec#Method_values),

>  . . . a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`.

You can see this in this example:

@file value_methods.go output sha1:22bffb5e9cb5efd7631839d4296d1203cb5fc568 goplayground:OuqIxPF10ef

To learn more about pointer and value methods, visit the [Go Spec section on Method Values](https://golang.org/ref/spec#Method_values), or see the [Effective Go section about Pointers v. Values](https://golang.org/doc/effective_go.html#pointers_vs_values).

_Note 1: The parenthesis (`()`) around `*p` and `&f` before selectors like `.Bar` are there for grouping purposes, and must be kept._

_Note 2: Although pointers can be converted to values (and vice-versa) when they are the receivers for a method, they are_ not _automatically converted to each other when they are arguments inside of a function._
