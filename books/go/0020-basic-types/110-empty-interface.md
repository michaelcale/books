Title: Empty interface
Id: 90100072
Body:
Type `interface{}` is called empty interface.

By definition it's an [interface](a-9010008c) with no methods defined.

That means that every type conforms to empty interface.

In Go it's used as a dynamic type (similar to object in C# or Java) because you can assign any value to a variable of `interface{}`.

[Zero value](a-6069) of empty interface is nil.
