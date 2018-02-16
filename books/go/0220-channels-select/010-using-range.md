---
Title: Using range to read from a channel
Id: 4134
---
When reading multiple values from a channel, using `range` is a common pattern:

@file range.go output sha1:19ae0b16e02794c207b8cd340299840413d5c499 goplayground:kM-SL7aRwXM

Using a `for range` loop is one of three ways to read values from a channel.

The loop ends when the channel is closed.

This is a common pattern when using worker pool:
* create a single channel for all workers
* launch workers
* workers use `for v := range chan` to pick up jobs to process
* after enquing all jobs, close the channel so that worker goroutines end when they process all jobs from a channel
