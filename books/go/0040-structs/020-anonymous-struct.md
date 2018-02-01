Title: Anonymous struct
Id: 1299
Score: 18
Body:
It is possible to create an anonymous struct:

    data := struct {
        Number int 
        Text   string
    } { 
        42, 
        "Hello world!",
    }

Full example:

    package main
    
    import (
        "fmt"
    )
    
    func main() {
        data := struct {Number int; Text string}{42, "Hello world!"} // anonymous struct
        fmt.Printf("%+v\n", data)
    }
[play it on playground](https://play.golang.org/p/atpNnP5wE_)
|======|
