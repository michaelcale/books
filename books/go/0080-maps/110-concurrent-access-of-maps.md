---
Title: Concurrent Access of Maps
Id: 3423
Score: 9
---
Maps in go are not safe for concurrency. You must take a lock to read and write on them if you will be accessing them concurrently. Usually the best option is to use `sync.RWMutex` because you can have read and write locks. However, a `sync.Mutex` could also be used.

@file concurrent_access.go

TODO: write a better example

TODO: talk about sync.Map