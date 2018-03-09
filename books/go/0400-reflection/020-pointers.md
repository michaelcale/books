---
Title: Pointers
Id: 226
SOId: 6073
---
A pointer to X is a distinct type from X.

If `reflect.Value` references a pointer to a value, you can get a reference to the value by calling `Elem()`.

@file pointers.go output

Under the hood an interface is also a pointer to its underlying value so `Elem()` also works on `reflect.Value` representing interface.

