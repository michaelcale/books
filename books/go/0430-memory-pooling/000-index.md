---
Title: Memory pooling
Id: 241
SOId: 4647
---
## Introduction
sync.Pool stores a cache of allocated but unused items for future use, avoiding memory churn for frequently changed collections, and allowing efficient, thread-safe re-use of memory. It is useful to manage a group of temporary items shared between concurrent clients of a package, for example a list of database connections or a list of output buffers.
