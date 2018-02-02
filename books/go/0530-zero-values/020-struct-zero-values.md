Title: Struct Zero Values
Id: 21171
Score: 1
Body:
When creating a struct without initializing it, each field of the struct is initialized to its respective zero value.

    type ZeroStruct struct {
        myString string
        myInt    int64
        myBool   bool
    }
    
    func main() {
        var myZero = ZeroStruct{}
        fmt.Printf("Zero values are: %q, %d, %t\n", myZero.myString, myZero.myInt, myZero.myBool)
        // Prints "Zero values are: "", 0, false"
    }
|======|
