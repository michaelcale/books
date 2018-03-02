---
Title: Writing cancellable functions
Id: 10386
---
Using existing functions that accept cancellable context is easy.

Writing a function that can be cancelled via context is much harder.

When a time experies or you call cancel function returned by `context.WithCancel()` or `context.WithTimeout()`, a channel in the context is signalled.

When writing a cancellable function, you have to periodically check channel returned by `context.Done()` and return immediately if it has been signalled.

It does make for an awkward code:

@file cancellable_function.go output sha1:29d550650156898244a024cf6258877b508481ed goplayground:AihF82Xoffg

For clarity, this is an artificial task.

Our `longMathOp` function performs simple operation 100 times and simulates slowness by sleeping for 1 ms on every iteration.

We can expect it to take ~100 ms.

A `select` with `default` clause is non-blocking. If there's nothing in the `ctx.Done()` channel, we don't wait for values and immediately execute `default` part, which is where the logic of the program lives.

We can see in our test that if timeout is greater that 100 ms, the function finishes.

If timeout is smaller than 100 ms, `ctx.Done()` channel is signalled, we detect it in `longMathOp` and return `ctx.Err()`.
