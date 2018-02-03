Title: Break-continue statements
Id: 23401
Score: 0
Body:
The break statement, on execution makes the current loop to force exit    

package main
    
    import "fmt"
    
    func main() {
        i:=0
        for true {
          if i>2 {
            break
            }
        fmt.Println("Iteration : ",i)
        i++
        }
    }

The continue statement, on execution moves the control to the start of the loop
    
    import "fmt"
    
    func main() {
        j:=100
        for j<110 {
         j++
         if j%2==0 {
            continue
            } 
        fmt.Println("Var : ",j)        
        }
    }

Break/continue loop inside switch

    import "fmt"
    
    func main() {
        j := 100
    
    loop:
        for j < 110 {
            j++
    
            switch j % 3 {
            case 0:
                continue loop
            case 1:
                break loop
            }
    
            fmt.Println("Var : ", j)
        }
    }


|======|
