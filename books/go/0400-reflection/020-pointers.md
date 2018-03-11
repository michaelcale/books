---
Title: Pointers
Id: 226
SOId: 6073
---
A pointer to X is a distinct type from X.

If `reflect.Value` references a pointer to a value, you can get a reference to the value by calling `Elem()`.

@file pointers.go output sha1:ebc16630ecdaaa6b3075c27154f193423e43df38 goplayground:aXx2H0KaebD

Under the hood an interface is also a pointer to its underlying value so `Elem()` also works on `reflect.Value` representing interface.

