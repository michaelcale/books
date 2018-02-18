---
Title: Concurrency
Id: 376
---
Go uses goroutines for concurrency. Simplifying, they are like threads.

Gorutines execute independently but they share memory space. In other words, they all see global variables.

To coordinate work between goroutines Go provides channels, which are thread-safe queues.

Here's an example of using worker pool of goroutines and coordinating their work with channels:

@file index.go output sha1:a4d1d5f37a28cb84a230fafab06e96384f1560f1 goplayground:7BD9mDiG0Uu

There's a lot to unpack here.

We launch 2 workers with `go sqrtWorker(chIn, chOut)`.

 Each `sqrtWorker` function is running independently and concurrently with all other code.

We use a single channel of int values to queue work items to be processed by worker goroutines using `<-` send operation on a channel.

Most of the time it's not safe to access the same variable from multiple goroutines. Channels and `sync.WaitGroup` are exceptions.

Worker goroutines `sqrtWorker` pick up values from the channel using `range`.

We don't know which worker will pick any given value.

Worker's `for` loop terminates when `chIn` is closed with `close(chIn)` and worker goroutine terminates.

To pass results of worker goroutines back to the caller we use another channel.

In this example we use unbuffered channels which only have capacity for one item at a time. For that reason we launch another goroutine to fill `chIn`. Otherwise we would risk dead-lock.

To shutdown workers we close the `chIn`.

We then wait for results created by workers by iterating on `chOut`.

There's one more complication. Unless we close `chOut`, the `for sqrt := range chOut` loop will wait forever.

To stop the loop, we need to `close(chOut)` but when to do it?

We can't do it in `sqrtWorker` because there are many of them and calling `close` on an already closed channel will [panic](ch-4350).

`sync.WaitGroup` is a thread-safe counter that can be incremented / decremented and allows for waiting until the counter reaches 0.

Before we launch the worker, we increment `wg` counter.

Just before terminating, the worker decrements `wg` counter.

We then `wg.Wait()` for the counter to reach 0, which indicates that all workers have finished and it's now safe to close the output channel.

It has to happen in its own goroutine to avoid blocking.

Why go to all this trouble to calculate a simple value?

This is just an example. In a real programs worker goroutine would perform longer jobs like downloading a file from the internet or resizing a large image.

