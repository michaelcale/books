---
Title: Recovering from panic
Id: 22031
---
A common mistake is to declare a slice and start requesting indexes from it without initializing it, which leads to an "index out of range" panic.

The following code explains how to recover from the panic without exiting the program, which is the normal behavior for a panic.

In most situations, returning an error in this fashion rather than exiting the program on a panic is only useful for development or testing purposes.

@file recover_from_panic.go

The use of a separate function (rather than closure) allows re-use of the same function in other functions prone to panic.
