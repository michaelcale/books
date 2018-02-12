---
Title: Chaining methods
Id: 13465
Score: 0
---
With methods in golang you can do method "chaining" passing pointer to method and returning pointer to the same struct like this:

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

func (empl *Employee) Promote() *Employee {
    fmt.Printf("Promoting %s\n", empl.Name)
    empl.Rank++
    return empl
}

func (empl *Employee) SetName(name string) *Employee {
    fmt.Printf("Set name of new Employee to %s\n", name)
    empl.Name = name
    return empl
}

func main() {

    worker := new(Employee)

    worker.Rank = 1

    worker.SetName("Bob").Promote()

    fmt.Printf("Here we have %s with rank %d\n", worker.Name, worker.Rank)
}
```

Output:

```text
Set name of new Employee to Bob
Promoting Bob
Here we have Bob with rank 2
```
