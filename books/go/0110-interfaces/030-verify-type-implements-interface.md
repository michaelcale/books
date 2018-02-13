---
Title: Testing Interface Implementation
Id: 9653
Score: 0
---

In Go interfaces are satisfied implicitly. You don't have to declare that a type is supposed to implement a given interface.

It's convenient but also makes it possible to not fully implement an interface by mistake and compiler has no way of detecting that.


There's a way to a compile-type check for that:

@file verify_type_implements_interface.go output allow_error

Our intent was for `MyReadCloser` to implement `io.ReadCloser` interface.

However, we forgot to implement `Close` method.

This line caught this problem at compile time:

```go
var _ io.ReadCloser = &MyReadCloser{}
```

We tried to assign `*MyReadCloser` type to variable of type `io.ReadCloser`.

Since `*MyReadCloser` doesn't implement `Close` method, the compiler detected this is an invalid assignement at compile time.

We assigned the value to [blank identifier](29103) `_` because we don't actually use that variable for anything.
