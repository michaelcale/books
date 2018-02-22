---
Title: Detecting races
Id: rd6000p1
---
When you don't use `sync.Mutex` to ensure exclusive access to data between goroutines or forget to lock in parts of the program, you'll get data races.

Data races might lead to memory corruption or crashes.

Go makes it easy to instrument the code with additional checks that are very likely to catch data races.

Use `-race` flag to `go build` or `go run`.

Here's a program with intentional data races.

When you run it with `go run -race data_race.go` the runtime will notice memory corruption.

@file data_race.go output allow_error sha1:86a300bf5d91df01d58fd9509f905e93dbef934d goplayground:fy_-Ws0cm98

This examples shows that memory for variable `n` is corrupted because the final value of `n` is not what we expect.

It also shows that instrumentation added with `-race` can catch memory corruption and points out which part of program caused the corruption.

## When to use `-race`

Additional instrumentation added by `-race` flag makes the program slower so it's not usually used when compiling shipping binaries.

It's a good idea to use on your CI (continous integration) servers when running your test suite.
