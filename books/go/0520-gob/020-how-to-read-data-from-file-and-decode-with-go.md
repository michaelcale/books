---
Title: How to read data from file and decode with go?
Id: 278
Score: 1
SOId: 27480
---
```go
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
```
