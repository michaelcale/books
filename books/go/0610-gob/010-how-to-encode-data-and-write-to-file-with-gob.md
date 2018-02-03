Title: How to encode data and write to file with gob?
Id: 27479
Score: 1
Body:
    package main

    import (
        "encoding/gob"
        "os"
    )

    type User struct {
        Username string
        Password string
    }

    func main() {

        user := User{
            "zola",
            "supersecretpassword",
        }

        file, _ := os.Create("user.gob")

        defer file.Close()

        encoder := gob.NewEncoder(file)

        encoder.Encode(user)

    }

|======|
