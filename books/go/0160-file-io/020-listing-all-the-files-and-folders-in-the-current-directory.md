Title: Listing all the files and folders in the current directory
Id: 3333
Score: 2
Body:
    package main
    
    import (
        "fmt"
        "io/ioutil"
    )
    
    func main() {
        files, err := ioutil.ReadDir(".")
        if err != nil {
            panic(err)
        }
    
        fmt.Println("Files and folders in the current directory:")
    
        for _, fileInfo := range files {
            fmt.Println(fileInfo.Name())
        }
    }

|======|
