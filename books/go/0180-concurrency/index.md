---
Title: Concurrency
Id: 376
---
## Introduction
In Go, concurrency is achieved through the use of goroutines, and communication between goroutines is usually done with channels. However, other means of synchronization, like mutexes and wait groups, are available, and should be used whenever they are more convenient than channels.

## Syntax
 - go doWork() // run the function doWork as a goroutine
 - ch := make(chan int) // declare new channel of type int
 - ch <- 1 // sending on a channel
 - value = <-ch // receiving from a channel

## Remarks
Goroutines in Go are similar to threads in other languages in terms of usage. Internally, Go creates a number of threads (specified by `GOMAXPROCS`) and then schedules the goroutines to run on the threads. Because of this design, Go's concurrency mechanisms are much more efficient than threads in terms of memory usage and initialization time.

