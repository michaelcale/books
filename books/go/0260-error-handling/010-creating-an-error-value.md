Title: Creating an error value
Id: 2705
Score: 12
Body:
The simplest way to create an error is by using the [`errors`](https://golang.org/pkg/errors/) package.

    errors.New("this is an error")

If you want to add additional information to an error, the [`fmt`](https://golang.org/pkg/fmt/) package also provides a useful error creation method:

    var f float64
    fmt.Errorf("error with some additional information: %g", f)

Here's a full example, where the error is returned from a function:

    package main

    import (
        "errors"
        "fmt"
    )

    var ErrThreeNotFound = errors.New("error 3 is not found")

    func main() {
        fmt.Println(DoSomething(1)) // succeeds! returns nil
        fmt.Println(DoSomething(2)) // returns a specific error message
        fmt.Println(DoSomething(3)) // returns an error variable
        fmt.Println(DoSomething(4)) // returns a simple error message
    }

    func DoSomething(someID int) error {
        switch someID {
        case 3:
            return ErrThreeNotFound
        case 2:
            return fmt.Errorf("this is an error with extra info: %d", someID)
        case 1:
            return nil
        }

        return errors.New("this is an error")
    }

[Open in Playground][1]


  [1]: https://play.golang.org/p/4xlwXJo2m0
|======|
