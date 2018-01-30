Title: How to read data from file and decode with go?
Id: 27480
Score: 1
Body:
    package main

    import (
        "encoding/gob"
        "fmt"
        "os"
    )

    type User struct {
        Username string
        Password string
    }

    func main() {

        user := User{}

        file, _ := os.Open("user.gob")

        defer file.Close()

        decoder := gob.NewDecoder(file)

        decoder.Decode(&user)

        fmt.Println(user)

    }

|======|
