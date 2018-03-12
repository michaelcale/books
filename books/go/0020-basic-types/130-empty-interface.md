---
Title: Empty interface
Id: 25
SOId: 90100072
---

Type `interface{}` is called an empty interface.

By definition it's an [interface](24) with no methods defined.

That means that every type conforms to an empty interface.

In Go `interface{}` is used as a dynamic type.

It's similar to an object in C# or Java because it combines a type and a value into a single value.

Empty interface is also how you implement a union type in Go.

[Zero value](29) of empty interface is nil.

Learn more about [empty interfaces](94).
