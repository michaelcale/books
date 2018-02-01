Title: Empty struct
Id: 20546
Score: 2
Body:
A struct is a sequence of named elements, called fields, each of which has a name and a type. Empty struct has no fields, like this anonymous empty struct:

    var s struct{}

Or like this named empty struct type:

    type T struct{}

The interesting thing about the empty struct is that, its size is zero (try [The Go Playground][2]):

    fmt.Println(unsafe.Sizeof(s))

This prints `0`, so the empty struct itself takes no memory. so it is good option for quit channel, like (try [The Go Playground][1]):

 

    package main
    
    import (
        "fmt"
        "time"
    )
    
    func main() {
        done := make(chan struct{})
        go func() {
            time.Sleep(1 * time.Second)
            close(done)
        }()
    
        fmt.Println("Wait...")
        <-done
        fmt.Println("done.")
    }


----------



  [1]: https://play.golang.org/p/j3qowmGdmC
  [2]: https://play.golang.org/p/ICQkZn01ng
|======|
