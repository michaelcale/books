---
Title: Buffered vs. unbuffered channels
Id: 20762
---
By default channels are unbuffered.

Sending and receiving goroutines block unless sending goroutine has a value to send and receiving goroutine is ready to receive.

Insisting on synchronization for every receive/send operation might introduce unnecessary slowdowns.

Imagine a scenario where one worker produces values and another worker consumes it.

If it takes a second to produce a value and a second to consume it, it takes 2 * N seconds to produce and consume all values.

If producer can queue up multiple values in the channel, producer doesn't have to wait for consumer to be ready for each value.

This is a job for bufferred channels.

By allowing producer to proceed independently of the consumer we can speed up some scenarios.

@file buffered.go output

