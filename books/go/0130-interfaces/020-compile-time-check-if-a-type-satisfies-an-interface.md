---
Title: Compile-time check if a type satisfies an interface
Id: 14631
Score: 6
---
Interfaces and implementations (types that implement an interface) are "detached". So it is a rightful question how to check at compile-time if a type implements an interface.

One way to ask the compiler to check that the type `T` implements the interface `I` is by attempting an assignment using the zero value for `T` or pointer to `T`, as appropriate. And we may choose to assign to the [blank identifier](https://golang.org/ref/spec#Blank_identifier) to avoid unnecessary garbage:

```go
type T struct{}

var _ I = T{}       // Verify that T implements I.
var _ I = (*T)(nil) // Verify that *T implements I.
```

If `T` or `*T` does not implement `I`, it will be a compile time error.

This question also appears in the official FAQ: [How can I guarantee my type satisfies an interface?](https://golang.org/doc/faq#guarantee_satisfies_interface)
