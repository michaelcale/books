---
Title: Empty interface
Id: der300hf
---
Technically speaking, an empty interface (`interface{}`) is an [interface](a-1221) with no methods.

What follows from that is that every type conforms to `interface{}`.

In practice, empty interface is Go's version of `object` type in Java or C# in that it combines a type and its value.

Empty interface is effectively a dynamic type in a static language.

Empty interface is also a way to implement union types in Go.

Since every type conforms to `interface{}`, you can assign any value to a variable of `interface{}` type.

At that point, you can no longer tell what is the real type at compile time.

You can query the type at runtime using:
* [type assertion](a-25362)
* [type switch](a-14736)

[Zero value](a-6069) of empty interface is nil.

Basic example:

@file index.go output sha1:d0df50dae2e229e7436bc7bfb0d158624fc1830b goplayground:ScwPsMsm6tT

<!-- TODO: how interface is implemented -->

<!-- TODO: describe a gotcha of nil vs. value is nil -->
