---
Title: Mutex gotchas
Search: Mutex pitfalls
Id: 801000u9
---
## Don't copy a mutex

A copy of `sync.Mutex` variable starts with the same state as original mutex but is not the same mutex.

It's almost always a mistake to copy a `sync.Mutex` e.g. by passing it to another function or embedding it in a struct and making a copy of that struct.

If you want to share a mutex variable, pass it around as a pointer.

## Mutex is not recursive

In some languages mutexes are recursive i.e. the same thread can `Lock` the same thread multiple times, as long as it calls `Unlock` the same number of times.

In Go `sync.Mutex` is non-recursive. If the same goroutine calls `Lock` twice on the same mutex, the program will deadlock.
