---
Title: Methods
Id: 1257
---
Struct methods are functions attached to structs:

@file methods.go output

[Playground](https://play.golang.org/p/I5e3yOaRcI)

The only difference is the addition of the method receiver.

It may be declared either as an instance of the type or a pointer to an instance of the type.

Since `SetName()` mutates the instance, the receiver must be a pointer in order to effect a permanent change in the instance.

<!-- TODO: more on pointer vs. value methods -->
