---
Title: Limiting concurrency with semaphore
Id: 901000od
---
When doing CPU intensive tasks like image resizing it doesn't make sense to create more goroutines than available CPUs.

Things will not go any faster and you might even loose performance due to additional book keeping and switching between goroutines.

One way to limit concurrency (i.e. number of gorutines you launch at the same time) is to use a semaphore.

You can enter (acquire) semaphore and leave (release) a semaphore.

A semaphore has a fixed capacity. If you exceed semaphore's capacity, acquire blocks until a release operation frees it.

A buferred channel is a natural semaphore.

Here's an example of using a channel acting as a semaphore to limit number of gouroutines active at any given time:

@file limit_with_semaphore.go output

We use technique described in [waiting for goroutines to finish](a-2490) to wait for all tasks to finish.

Often the right amount of concurrent tasks is equal to number of CPUs, which can be obtained with [`runtime.NumCPU()`](https://golang.org/pkg/runtime/#NumCPU).

