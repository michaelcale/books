---
Title: Defer basics
Id: 129
SOId: 9429
---
`defer` statements marks a function to be executed at a later time.

Defer statement is an ordinary function call prefixed by the keyword `defer`.

```go
defer someFunction()
```

A deferred function is executed once the function that contains the `defer` statement returns. Actual call to the deferred function occurs when the enclosing function:
- executes a return statement
- falls off the end
- panics

Example:

@file defer_basics.go output sha1:be9a912a829d01503ae29216159d5fc776fef713 goplayground:qDNfakDI7g6

If a function has multiple deferred statements, they form a stack. The last `defer` is the first one to execute after the enclosing function returns, followed by subsequent calls to preceding `defer`s in order (below example returns by causing a panic):

@file defer_basics2.go output allow_error sha1:dceece5d882d36b367ab82895afef4dd1f1e034c goplayground:DSIsTuR7ifw

Note that deferred functions have their arguments evaluated at the time `defer` executes:

@file defer_basics3.go output sha1:c322cb89e29edc888998f53e0192537879bc6565 goplayground:cdlR5oUXrAs

If a function has named return values, a deferred anonymous function within that function can access and update the returned value even after the function has returned:

@file defer_basics4.go output sha1:25e79d5067349dde1475059fac942d1e34977098 goplayground:FHb-GI09tOV

Finally, a `defer` statement is generally used operations that often occur together. For example:
- open and close a file
- connect and disconnect
- lock and unlock a mutex
- mark a waitgroup as done (`defer wg.Done()`)

This use ensures proper release of system resources irrespective of the flow of execution.

```go
resp, err := http.Get(url)
if err != nil {
return err
}
defer resp.Body.Close() // Body will always get closed
```
