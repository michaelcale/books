---
Title: Type Assertion
Id: 25362
Score: 0
---
You can access the real data type of interface with Type Assertion.

```go
interfaceVariable.(DataType)
```

Example of struct `MyType` which implement interface `Subber`:

```go
package main

import (
    "fmt"
)

type Subber interface {
    Sub(a, b int) int
}

type MyType struct {
    Msg string
}

//Implement method Sub(a,b int) int
func (m *MyType) Sub(a, b int) int {
    m.Msg = "SUB!!!"

    return a - b;
}

func main() {
    var interfaceVar Subber = &MyType{}
    fmt.Println(interfaceVar.Sub(6,5))
    fmt.Println(interfaceVar.(*MyType).Msg)
}
```

Without `.(*MyType)` we wouldn't able to access `Msg` Field. If we try `interfaceVar.Msg` it will show compile error:

```go
interfaceVar.Msg undefined (type Subber has no field or method Msg)
```
