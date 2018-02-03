Title: Decoding JSON data from a file
Id: 6628
Score: 1
Body:
JSON data can also be read from files.

Let's assume we have a file called `data.json` with the following content:


    [
        {
          "Name" : "John Doe",
          "Standard" : 4
        },
        {
          "Name" : "Peter Parker",
          "Standard" : 11
        },
        {
          "Name" : "Bilbo Baggins",
          "Standard" : 150
        }
    ]

The following example reads the file and decodes the content:

    package main
    
    import (
        "encoding/json"
        "fmt"
        "log"
        "os"
    )
    
    type Student struct {
        Name     string
        Standard int `json:"Standard"`
    }
    
    func main() {
        // open the file pointer
        studentFile, err := os.Open("data.json")
        if err != nil {
            log.Fatal(err)
        }
        defer studentFile.Close()

        // create a new decoder
        var studentDecoder *json.Decoder = json.NewDecoder(studentFile)
        if err != nil {
            log.Fatal(err)
        }

        // initialize the storage for the decoded data
        var studentList []Student
        
        // decode the data
        err = studentDecoder.Decode(&studentList)
        if err != nil {
            log.Fatal(err)
        }

        for i, student := range studentList {
            fmt.Println("Student", i+1)
            fmt.Println("Student name:", student.Name)
            fmt.Println("Student standard:", student.Standard)
        }
    }

The file `data.json` must be in the same directory of the Go executable program. Read [Go File I/O documentation][1] for more information on how to work with files in Go.

  [1]: http://stackoverflow.com/documentation/go/1033/file-i-o

|======|
