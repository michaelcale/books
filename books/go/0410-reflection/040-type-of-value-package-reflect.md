Title: Type of value - package "reflect"
Id: 23400
Score: 0
Body:
reflect.TypeOf  can be used to check the type of variables when comparing
    
    package main
        
        import (
            "fmt"
            "reflect"
        )
        type Data struct {
         a int
        }
        func main() {
            s:="hey dude"
            fmt.Println(reflect.TypeOf(s))
            
            D := Data{a:5}
            fmt.Println(reflect.TypeOf(D))
            
        }

Output :  
string  
main.Data


|======|
