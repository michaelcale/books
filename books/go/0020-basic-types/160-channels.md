---
Title: Channels
Id: 901000c2
---
Channels are typed queues used for goroutines to communicate with each other in a thread-safe manner.

Channel basics:
```go
// create unbuffered channel of int values with capacity of 1
ch := make(chan int)
// start a new goroutine that sends value 3 over a channel
go func() { ch <- 3 }()
// read the value from a channel
// it waits until goroutine above sends a value
n := <-ch
fmt.Printf("n: %d\n", n)
```

[Zero value](a-6069) of a channel is nil.

Learn more about [channels](ch-1263).