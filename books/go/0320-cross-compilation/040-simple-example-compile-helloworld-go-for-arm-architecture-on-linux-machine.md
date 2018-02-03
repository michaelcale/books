Title:
Simple Example: Compile helloworld.go for arm architecture on Linux machine
|======|
Id: 31164
Score: 0
Body:
**Prepare** helloworld.go (find below)

    package main
    
    import "fmt"
    
    func main(){
            fmt.Println("hello world")
    }

**Run** `GOOS=linux GOARCH=arm go build helloworld.go`

**Copy** generated `helloworld` (arm executable) file to your target machine.


 

|======|
