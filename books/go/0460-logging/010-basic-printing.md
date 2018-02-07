---
Title: Basic Printing
Id: 12872
Score: 1
---
Go has a built-in logging library known as `log` with a commonly use method `Print` and its variants. You can import the library then do some basic printing:

```go
package main

import "log"

func main() {

    log.Println("Hello, world!")
    // Prints 'Hello, world!' on a single line

    log.Print("Hello, again! \n")
    // Prints 'Hello, again!' but doesn't break at the end without \n

    hello := "Hello, Stackers!"
    log.Printf("The type of hello is: %T \n", hello)
    // Allows you to use standard string formatting. Prints the type 'string' for %T
    // 'The type of hello is: string
}
```
