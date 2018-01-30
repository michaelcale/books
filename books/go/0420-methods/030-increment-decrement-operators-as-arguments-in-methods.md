Title: Increment-Decrement operators as arguments in Methods
Id: 23512
Score: 0
Body:
Though Go supports ++ and -- operators and the behaviour is found to be almost similar to c/c++, variables with such operators cannot be passed as argument to function.

        package main
    
        import (
            "fmt"
        )
        
        func abcd(a int, b int) {
         fmt.Println(a," ",b)
        }
        func main() {
            a:=5
            abcd(a++,++a)
        }

Output: syntax error: unexpected ++, expecting comma or )


|======|
