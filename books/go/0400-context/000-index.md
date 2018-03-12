---
Title: Context
Id: 231
SOId: 2743
---

Package `context` in standard library provides type `Context` which is hard to explain because it has multiple uses.

Here are the most common uses of `context.Context`:

* context with timeout (deadline) is a generic way to implement timeouts for functions that can possibly take a long time where we want an option to abort them if they exceed the timeout
* context with cancellation is a generic way to cancel a goroutine
* context with value is a way to associate arbitrary value with a context

## Creating a context

In most cases you'll be calling existing API that requires `context.Context`.

If you don't have one, use `context.TODO()` or `context.Background()` functions to create it. Read about [the difference](235).

`context.Context` is an immutable (read-only) value so you can't modify it.

To create e.g. a context with value, you call `context.WithValue()` which returns a new context that wraps existing context and adds additional information.

`context.Context` is an interface so you could pass `nil` but it's not recommended.

Many APIs expect non-nil value and will crash if passed `nil` so it's best to always pass one created with `context.Background()` or `context.TODO()`.

There is no performance issue because those functions return shared, global variables (yay immutability!).

## Using context with timeout to set timeout for HTTP requests

Repeating an example from [HTTP client article](202) here's a way to create a context with timeout to ensure HTTP GET request doesn't hang forever:

@file ../0360-http-client/http_timeout.go output allow_error noplayground

HTTP client knows how to interpret context with a timeout. You just need to create and provide it.
