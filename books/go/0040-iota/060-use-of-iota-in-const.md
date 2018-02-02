Title: Use of iota in const
Id: 29627
Score: 1
Body:
This is an enumeration for const creation. Go compiler starts iota from 0 and increments by one for each following constant.  The value is determined at compile time rather than run time. Because of this we can't apply iota to expressions which are evaluated at run time. 

Program to use iota in const

    package main
    
    import "fmt"
    
    const (
        Low = 5 * iota
        Medium
        High
    )
    
    func main() {
        // Use our iota constants.
        fmt.Println(Low)
        fmt.Println(Medium)
        fmt.Println(High)
    }

Try it in [Go Playground][1]


  [1]: https://play.golang.org/p/jyJEzyZSi6
|======|
