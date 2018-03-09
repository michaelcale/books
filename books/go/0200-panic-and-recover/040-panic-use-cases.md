---
Title: Panic use cases
Id: 135
SOId: rd6000v3
---

Using of `panic` in Go should be extremely rare.

But it does have its uses.

## Force crash on bad API usage

This is similar to using `assert` in other languages.

Imagine you're writing function `sqrt(n int) int` that returns square root of a number.

Mathematically, that only makes sense for `n >= 0`.

What should happen if the function is called with `n < 0`?

You could change the function to also return an error.

Or you could force program to crash, under the assumption that calling `sqrt` with negative `n` indicates a bug in the calling code and that bug should be fixed.

```go
func sqrt(n int) int {
    if n < 0 {
        panic("sqrt only accepts n >= 0")
    }
    // ... calculate sqrt
}
```

## Simplifying flow control in isolated piece of code

Sometimes it's easier to propagate an error by panicking and recovering rather than returning error code up the call chain.

This might be true in e.g. parser.

When you use that technique, you should observe one rule: `panic` should not cross public API boundary.

In other words, if you implement a parser library, panic / recover should happen only within your package.

You should never create public API that expects the caller to `recover` a `panic` thrown in your code.

## Protect a program from a crash in a gouroutine

This technique is used in Go's http server package.

Each connection is handled in a separate routine.

A bug that causes a panic in one code path would crash the whole program.

This is not a good behavior for an http server so each goroutine is wrapped in a function that recovers all panics.
