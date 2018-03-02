---
Title: Context with value
Id: 9010008e
---
In HTTP server each request is served by a handler function running in its own goroutine.

We often want to have common per-requeste information available in a convenient way.

For example, at the beginning of handling a request we might check a cookie to see if a request is made by a logged in user and we want to have user info available everywhere.

We can do that by using context with value:

@file context_with_value.go output sha1:8d8a6526a597b72b6dd29e4b1700245eb6f26b9d goplayground:EkdcW_DuUNf

For clarity of the example we only show creating a context with value and retrieving value from context.

Because context value is an `interface{}`, it's a good practice to write type-safe wrapper functions for setting and retrieving values.

The key used to set / get value is also an `interface{}`. Because context can be passed to functions in code you didn't write, you want to ensure that the value used for key is unique.

That's why we defined a non-exported type `userKeyType` and used a non-exported global variable `userKey` of that type.

This ensures that code outside of our package can't possibly use this key.

This wouldn't be true if the key was e.g. a string (or any type available to multiple packages).

We wrote two functions for retrieving the value.

One panics when value is not set, another one returns `nil`.

Which one to use is a policy decision for your code.

Sometimes missing a variable on context means a bug in your program and you should use `mustGetUserFromContext` variant which panics in that case.

Sometimes missing a variable is expected and you can use `getUserFromContext` variant.
