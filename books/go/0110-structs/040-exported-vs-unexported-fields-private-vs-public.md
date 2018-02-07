---
Title: Exported vs. Unexported Fields (Private vs Public)
Id: 1255
Score: 10
---
Struct fields whose names begin with an uppercase letter are exported. All other names are unexported.

```go
type Account struct {
    UserID      int    // exported
    accessToken string // unexported
}
```

Unexported fields can only be accessed by code within the same package. As such, if you are ever accessing a field from a _different_ package, its name needs to start with an uppercase letter.

```go
package main

import "bank"

func main() {
    var x = &bank.Account{
        UserID: 1,          // this works fine
        accessToken: "one", // this does not work, since accessToken is unexported
    }
}
```

However, from within the `bank` package, you can access both UserId and accessToken without issue.

The package `bank` could be implemented like this:

```go
package bank

type Account struct {
    UserID int
    accessToken string
}

func ProcessUser(u *Account) {
    u.accessToken = doSomething(u) // ProcessUser() can access u.accessToken because
                                    // it's defined in the same package
}
```
