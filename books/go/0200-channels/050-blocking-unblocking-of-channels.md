---
Title: Blocking & unblocking of channels
Id: 25771
Score: 0
---
By default communication over the channcels is sync; when you send some value there must be a receiver. Otherwise you will get `fatal error: all goroutines are asleep - deadlock!` as follows:

```go
package main

import "fmt"

func main() {
    msg := make(chan string)
    msg <- "Hey There"
    go func() {
        fmt.Println(<-msg)
    }()
}
```
But there is a solution use: use buffered channels :

```go
package main

import "fmt"
import "time"

func main() {
    msg :=make(chan string, 1)
    msg <- "Hey There!"
    go func() {
        fmt.Println(<-msg)
    }()
    time.Sleep(time.Second * 1)
}
```

