---
Title: Typed vs. Untyped Constants
Id: 12431
---
Constants in Go may be typed or untyped. For instance, given the following string literal:

    "bar"

one might say that the type of the literal is `string`, however, this is not semantically correct. Instead, literals are *Untyped string constants*. It is a string (more correctly, its *default type* is `string`), but it is not a Go **value** and therefore has no type until it is assigned or used in a context that is typed. This is a subtle distinction, but a useful one to understand.

Similarly, if we assign the literal to a constant:

```go
    const foo = "bar"
```

It remains untyped since, by default, constants are untyped. It is possible to declare it as a *typed string constant* as well:

```go
    const typedFoo string = "bar"
```

The difference comes into play when we attempt to assign these constants in a context that does have type. For instance, consider the following:

```go
var s string
s = foo      // This works just fine
s = typedFoo // As does this

type MyString string
var mys MyString
mys = foo      // This works just fine
mys = typedFoo // cannot use typedFoo (type string) as type MyString in assignment
```
