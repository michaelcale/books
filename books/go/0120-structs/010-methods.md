---
Title: Methods
Id: 1257
Score: 18
---
Struct methods are very similar to functions:

```go
type User struct {
    name string
}

func (u User) Name() string {
    return u.name
}

func (u *User) SetName(newName string) {
    u.name = newName
}
```

The only difference is the addition of the method receiver. It may be declared either as an instance of the type or a pointer to an instance of the type. Since `SetName()` mutates the instance, the receiver must be a pointer in order to effect a permanent change in the instance.

For example:

```go
package main

import "fmt"

type User struct {
    name string
}

func (u User) Name() string {
    return u.name
}

func (u *User) SetName(newName string) {
    u.name = newName
}

func main() {
    var me User

    me.SetName("Slim Shady")
    fmt.Println("My name is", me.Name())
}
```

[Go Playground](https://play.golang.org/p/I5e3yOaRcI)

