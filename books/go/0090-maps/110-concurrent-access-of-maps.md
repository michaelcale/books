---
Title: Concurrent access of maps
Id: 3423
---
Maps in Go are not safe for concurrency. You must take a lock to read and write to them if you will be accessing them concurrently. Usually the best option is to use `sync.RWMutex` because you can have read and write locks. However, a `sync.Mutex` could also be used.

@file concurrent_access.go sha1:41dc31b1fd554d4350995ef33153ac361a7085e8 goplayground:wetl0PWoZk_V

TODO: write a better example

TODO: talk about sync.Map