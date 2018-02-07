---
Title: Assigning signals to a channel
Id: 15744
Score: 1
---
Often times you will have reason to catch when your program is being told to stop by the OS and take some actions to preserve the state, or clean up your application. To accomplish this you can use the `os/signal` package from the standard library. Below is a simple example of assigning all signals from the system to a channel, and then how to react to those signals.

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
)

func main() {
    // create a channel for os.Signal
    sigChan := make(chan os.Signal)

    // assign all signal notifications to the channel
    signal.Notify(sigChan)

    // blocks until you get a signal from the OS
    select {
    // when a signal is received
    case sig := <-sigChan:
        // print this line telling us which signal was seen
        fmt.Println("Received signal from OS:", sig)
    }
}
```

When you run the above script it will create a channel, and then block until that channel receives a signal.

```sh
$ go run signals.go
^CReceived signal from OS: interrupt
```
The `^C` above is the keyboard command `CTRL+C` which sends the `SIGINT` signal.
