Title: FizzBuzz
Id: 4593
Score: 36
Body:
Another example of "Hello World" style programs is [FizzBuzz][1]. This is one example of a FizzBuzz implementation. Very idiomatic Go in play here.

    package main

    // Simple fizzbuzz implementation

    import "fmt"
    
    func main() {
        for i := 1; i <= 100; i++ {
            s := ""       
            if i % 3 == 0 {
                s += "Fizz"
            }
            if i % 5 == 0 {
                s += "Buzz"
            }
            if s != "" {
                fmt.Println(s) 
            } else {
                fmt.Println(i) 
            }
        }
    }

[Playground](https://play.golang.org/p/ckp5s9Fepm)

  [1]: https://blog.codinghorror.com/why-cant-programmers-program/
|======|
