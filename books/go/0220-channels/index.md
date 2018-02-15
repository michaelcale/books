---
Title: Channels
Id: 1263
---

A channel is a thread-safe queue of values of a given type.

A primary use for channels is to communicate between goroutines.

For example,

## Introduction

A channel contains values of a given type. Values can be written to a channel and read from it, and they circulate inside the channel in first-in-first-out order. There is a distinction between buffered channels, which can contain several messages, and unbuffered channels, which cannot. Channels are typically used to communicate between goroutines, but are also useful in other circumstances.

## Syntax
- make(chan int) // create an unbuffered channel
- make(chan int, 5) // create a buffered channel with a capacity of 5
- close(ch) // closes a channel "ch"
- ch <- 1 // write the value of 1 to a channel "ch"
- val := <-ch // read a value from channel "ch"
- val, ok := <-ch // alternate syntax; ok is a bool indicating if the channel is closed

## Remarks
A channel holding the empty struct `make(chan struct{})` is a clear message to the user that no information is transmitted over the channel and that it's purely used for synchronization.

Regarding unbuffered channels, a channel write will block until a corresponding read occurs from another goroutine. The same is true for a channel read blocking while waiting for a writer.
