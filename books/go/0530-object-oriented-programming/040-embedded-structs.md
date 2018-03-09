---
Title: Embedded structs
Id: 285
Score: 0
SOId: 27426
---
because a struct is also a data type, it can be used as an anonymous field, the outer struct can directly access the fields of the embedded struct even if the struct came from a diffrent package. this behaviour provides a way to derive some or all of your implementation from another type or a set of types.

```go
package main

type Admin struct {
    Username, Password string
}

type User struct {
    ID uint64
    FullName, Email string
    Admin // embedded struct
}

func main() {
    admin := Admin{
        "zola",
        "supersecretpassword",
    }

    user := User{
        1,
        "Zelalem Mekonen",
        "zola.mk.27@gmail.com",
        admin,
    }

    fmt.Println(admin) // {zola supersecretpassword}

    fmt.Println(user) // {1 Zelalem Mekonen zola.mk.27@gmail.com {zola supersecretpassword}}

    fmt.Println(user.Username) // zola

    fmt.Println(user.Password) // supersecretpassword
}
```
