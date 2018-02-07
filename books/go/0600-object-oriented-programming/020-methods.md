---
Title: Methods
Id: 27427
Score: 1
---
In Go a method is

> a function that acts on a variable of a certain type, called the receiver

the receiver can be anything, not only `structs` but even a `function`, alias types for built in types such as `int`, `string`, `bool` can have a method, an exception to this rule is that `interfaces`(discussed later) cannot have methods, since an interface is an abstract definition and a method is an implementation, trying it generate a compile error.

combining `structs` and `methods` you can get a close eqivalent of a `class` in Object Oriented programming.

a method in Go has the following signature

`func (name receiverType) methodName(paramterList) (returnList) {}`

```go
package main

type Admin struct {
    Username, Password string
}

func (admin Admin) Delete() {
    fmt.Println("Admin Deleted")
}

type User struct {
    ID uint64
    FullName, Email string
    Admin
}

func (user User) SendEmail(email string) {
    fmt.Printf("Email sent to: %s\n", user.Email)
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

    user.SendEmail("Hello") // Email sent to: zola.mk.27@gmail.com

    admin.Delete() // Admin Deleted
}
```
