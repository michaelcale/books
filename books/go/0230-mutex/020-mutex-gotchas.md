---
Title: Mutex gotchas
SearcH: Mutex pitfalls
Id: 801000u9
---
A copy of `sync.Mutex` variable starts with the same state as original mutex but is not the same mutex.

It's almost always a mistake to copy a `sync.Mutex` e.g. by passing it to another function or embedding it in a struct and making a copy of that struct.

If you want to share a mutex variable, pass it around as a pointer to `sync.Mutex`.
