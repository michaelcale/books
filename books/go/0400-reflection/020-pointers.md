---
Title: Pointers
Id: 6073
---
A pointer to a type X is a distinct type from X.

If `reflect.Value` references a pointer to a value, you can get a reference to the value by calling `Elem()`.

Under the hood an interface is also a pointer to its underlying value so `Elem()` also works on `reflect.Value` representing interface.

@file pointers.go output

There's no limit to nesting i.e. you can have a pointer to a pointer to a pointer...
