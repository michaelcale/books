---
Title: Read-Write mutes (RWMutex)
Id: 150
SOId: 8652
---
In a `sync.Mutex` `Lock()` always takes an exclusive lock.

In read-heavy scenarios we can improve performance if we allow multiple readers but only one writer.

A `sync.RWMutex` has 2 types of lock function: lock for reading and lock for writing.

It follows the following rules:
* a writer lock takes exclusive lock
* a reader lock will allow another readers but not writer

Here's a cache variant that uses read-write lock:

@file rwlock.go output sha1:ca40aabb9f1852ae977bc6020a13ade42d4d0b75 goplayground:HEsolwLu-0d
