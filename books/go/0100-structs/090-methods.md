---
Title: Methods
Id: 1257
---
Struct methods are functions attached to structs:

@file methods.go output sha1:aff9579292ce1e92613e9fd085e90f71e6278acc goplayground:xp3FjKhUefO

The only difference is the addition of the method receiver.

It may be declared either as an instance of the type or a pointer to an instance of the type.

Since `SetName()` mutates the instance, the receiver must be a pointer in order to effect a permanent change in the instance.

<!-- TODO: more on pointer vs. value methods -->
