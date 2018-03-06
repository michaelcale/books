---
Title: Reflection
Id: 1854
---
Go is a statically typed language. In most cases the type of a variable is known at compilation time.

One exception is interface type, especislly empty interface `interface{}`.

Empty interface is a dynamic type, similar to `Object` in Java or C#.

It hides the underlying real type at compile time.

Package `reflect` in standard library allows us to operate on dynamic values at runtime. You can:
* inspect the type of dynamic value
* enumerate fields of a struct
* set the value
* create new values at runtime

Related language-level functionality for inspecting type of an interface is a [type switch](a-14736) and a [type assertion](a-25362).

@file basic_reflect.go output
