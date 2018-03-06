---
Title: Reflection
Id: 1854
---
Go is a statically typed language. In most cases the type of a variable is known at compilation time.

One exception is interface type, especially empty interface `interface{}`.

Empty interface is a dynamic type, similar to `Object` in Java or C#.

At compilation time we can't tell if the underlying value of interface type is an `int` or a `string`.

Package `reflect` in standard library allows us to work with such dynamic values at runtime. We can:
* inspect the type of dynamic value
* enumerate fields of a struct
* set the value
* create new values at runtime

Related language-level functionality for inspecting type of an interface value at runtime is a [type switch](a-14736) and a [type assertion](a-25362).

@file basic_reflect.go output
