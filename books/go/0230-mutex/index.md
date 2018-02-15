---
Title: Mutex
Id: 2607
---
In Go all goroutines share memory.

This is good for performance but modyfing the same memory from multiple goroutines is not safe. It can read to data races and crashes.

One way to avoid that is by using [channels](ch-1263) to hand exclusive ownership of data.

This is Go's motto: *Do not communicate by sharing memory; instead, share memory by communicating.*

Another way to ensure exclusive access to data by goroutines is to use mutexes.

Here's a basic pattern that uses a map for a cache and ensure exclusive access to a map by locking.

@file index.go output

Notice that a zero-value of `sync.Mutex` is a valid mutex. You don't need to initialize a mutex.

For performance we want to minimize the time spent holding a lock.

Function `getCached` would be simpler if we kept the mutex locked for the duration of the function but we don't want to keep cache locked when we're executing `expensiveOperation()`.

Unlike many other languages, Go mutexes are non-recursive. If the same goroutines tries to `Lock()`` a mutex twice, the second `Lock()` will block forever.
