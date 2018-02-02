Title: Reading and writing to a file using ioutil
Id: 3332
Score: 4
Body:
A simple program that writes "Hello, world!" to `test.txt`, reads back the data, and prints it out. Demonstrates simple file I/O operations.

    package main
    
    import (
        "fmt"
        "io/ioutil"
    )
    
    func main() {
        hello := []byte("Hello, world!")
    
        // Write `Hello, world!` to test.txt that can read/written by user and read by others 
        err := ioutil.WriteFile("test.txt", hello, 0644)
        if err != nil {
            panic(err)
        }
    
        // Read test.txt
        data, err := ioutil.ReadFile("test.txt")
        if err != nil {
            panic(err)
        }
    
        // Should output: `The file contains: Hello, world!`
        fmt.Println("The file contains: " + string(data))
    }
|======|
