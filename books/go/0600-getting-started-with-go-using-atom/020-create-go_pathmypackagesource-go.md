Title: Create $GO_PATH/mypackage/source.go
Id: 26865
Score: 1
Body:
    package mypackage
    
    var PublicVar string = "Hello, dear reader!"
    
    //Calculates the factorial of given number recursively!
    func Factorial(x uint) uint {
        if x == 0 {
            return 1
        }
        return x * Factorial(x-1)
    }
    
|======|
