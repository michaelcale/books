---
Title: Basic methods
Id: 13464
Score: 0
---
Methods in Go are just like functions, except they have *receiver*.

Usually receiver is some kind of struct or type.

```go
package main

import (
    "fmt"
)

type Employee struct {
    Name string
    Age  int
    Rank int
}

func (empl *Employee) Promote() {
    empl.Rank++
}

func main() {

    Bob := new(Employee)

    Bob.Rank = 1
    fmt.Println("Bobs rank now is: ", Bob.Rank)
    fmt.Println("Lets promote Bob!")

    Bob.Promote()

    fmt.Println("Now Bobs rank is: ", Bob.Rank)
}
```

Output:

```text
Bobs rank now is:  1
Lets promote Bob!
Now Bobs rank is:  2
```
