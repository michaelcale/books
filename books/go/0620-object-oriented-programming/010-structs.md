Title: Structs
Id: 27425
Score: 1
Body:
Go supports user defined types in the form of structs and type aliases. structs are composite types, the component pieces of data that constitute the struct type are called *fields*. a field has a type and a name which must be unqiue.

    package main

    type User struct {
        ID uint64
        FullName string
        Email    string
    }

    func main() {
        user := User{
            1,
            "Zelalem Mekonen",
            "zola.mk.27@gmail.com",
        }

        fmt.Println(user) // {1 Zelalem Mekonen zola.mk.27@gmail.com}
    }

this is also a legal syntax for definining structs

    type User struct {
        ID uint64
        FullName, Email string
    }

    user := new(User)

    user.ID = 1
    user.FullName = "Zelalem Mekonen"
    user.Email = "zola.mk.27@gmail.com"
|======|
