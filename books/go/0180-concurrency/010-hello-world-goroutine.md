---
Title: Hello World Goroutine
Id: 1259
Score: 15
---
Single channel, single goroutine, one write, one read.

```go
package main

import "fmt"
import "time"

func main() {
    // create new channel of type string
    ch := make(chan string)

    // start new anonymous goroutine
    go func() {
        time.Sleep(time.Second)
        // send "Hello World" to channel
        ch <- "Hello World"
    }()
    // read from channel
    msg, ok := <-ch
    fmt.Printf("msg='%s', ok='%v'\n", msg, ok)
}
```

[Run it on playground](https://play.golang.org/p/t-5U31vPcb)

The channel `ch` is an **[unbuffered or synchronous channel](https://golang.org/doc/effective_go.html#channels)**.

The `time.Sleep` is here to illustrate `main()` function will **wait** on the `ch` channel, which means the [function literal](https://golang.org/ref/spec#Function_literals) executed as a goroutine has the time to send a value through that channel: the [receive operator `<-ch`](https://golang.org/ref/spec#Receive_operator) will block the execution of `main()`. If it didn't, the goroutine would be killed when `main()` exits, and would not have time to send its value.
