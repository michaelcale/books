---
Title: Stringer
Id: 237
Score: 0
SOId: 9983
---
The `fmt.Stringer` interface requires a single method, `String() string` to be satisfied. The string method defines the "native" string format for that value, and is the default representation if the value is provided to any of the `fmt` packages formatting or printing routines.

```go
package main

import (
    "fmt"
)

type User struct {
    Name  string
    Email string
}

// String satisfies the fmt.Stringer interface for the User type
func (u User) String() string {
    return fmt.Sprintf("%s <%s>", u.Name, u.Email)
}

func main() {
    u := User{
        Name:  "John Doe",
        Email: "johndoe@example.com",
    }

    fmt.Println(u)
    // output: John Doe <johndoe@example.com>
}
```
