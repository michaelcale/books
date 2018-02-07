---
Title: Goroutines Basic Program
Id: 30109
Score: 1
---
```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

A goroutine is a function that is capable of running concurrently with other functions. To create a goroutine we use the keyword **`go`** followed by a function invocation:

```go
package main

import "fmt"

func f(n int) {
    for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    }
}

func main() {
    go f(0)
    var input string
    fmt.Scanln(&input)
}
```

Generally, function call executes all the statements inside the function body and return to the next line. But, with goroutines we return immediately to the next line as it don't wait for the function to complete. So, a call to a `Scanln` function included, otherwise the program has been exited without printing the numbers.

