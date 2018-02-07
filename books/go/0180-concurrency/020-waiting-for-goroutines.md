---
Title: Waiting for goroutines
Id: 2490
Score: 11
---

[Go programs end when the `main` function ends][1], therefore it is common practice to wait for all goroutines to finish. A common solution for this is to use a [sync.WaitGroup][2] object.

```go
package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup // 1

func routine(i int) {
    defer wg.Done() // 3
    fmt.Printf("routine %v finished\n", i)
}

func main() {
    wg.Add(10) // 2
    for i := 0; i < 10; i++ {
        go routine(i) // *
    }
    wg.Wait() // 4
    fmt.Println("main finished")
}
```

[Run the example in the playground](https://play.golang.org/p/64vfZSXXHv)

WaitGroup usage in order of execution:

 1. Declaration of global variable. Making it global is the easiest way to make it visible to all functions and methods.
 2. Increasing the counter. This must be done in the main goroutine because there is no guarantee that a newly started goroutine will execute before 4 due to memory model [guarantees][3].
 3. Decreasing the counter. This must be done at the exit of a goroutine. By using a deferred call, we make sure that it [will be called whenever function ends][4], no matter how it ends.
 4. Waiting for the counter to reach 0. This must be done in the main goroutine to prevent the program from exiting before all goroutines have finished.

\* Parameters are [evaluated before starting a new goroutine][5]. Thus it is necessary to define their values explicitly before `wg.Add(10)` so that possibly-panicking code will not increase the counter. Adding 10 items to the WaitGroup, so it will wait for 10 items before `wg.Wait` returns the control back to `main()` goroutine. Here, the value of i is defined in the for loop.

  [1]: http://golang.org/ref/spec#Program_execution
  [2]: http://golang.org/pkg/sync/#WaitGroup
  [3]: http://golang.org/ref/mem#tmp_5
  [4]: http://golang.org/ref/spec#Defer_statements
  [5]: http://golang.org/ref/spec#Go_statements

