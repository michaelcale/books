---
Title: Panic and recover
Search: exception handling
Id: 4350
---
This chapter assumes knowledge of [defer](ch-2795).

In Go `panic` and `recover` are technically similar to exception handling in languages like C#, Java or Python.

`panic` is equivalent of `throw` or `raise` and `recover` fills the role of `catch`.

However, while those other languages often use exception handling for flow control, as a way to propagate errors up the call chain, using `panic` in Go is highly discouraged.

It should only be used in [truly exceptional cases](a-rd6000v3).

For ordinary error handling, read chapter on [error handling](ch-785).
