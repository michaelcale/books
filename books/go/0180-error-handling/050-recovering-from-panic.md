---
Title: Recovering from panic
Id: 22031
Score: 1
---
A common mistake is to declare a slice and start requesting indexes from it without initializing it, which leads to an "index out of range" panic. The following code explains how to recover from the panic without exiting the program, which is the normal behavior for a panic. In most situations, returning an error in this fashion rather than exiting the program on a panic is only useful for development or testing purposes.

```go
type Foo struct {
    Is []int
}

func main() {
    fp := &Foo{}
    if err := fp.Panic(); err != nil {
        fmt.Printf("Error: %v", err)
    }
    fmt.Println("ok")
}

func (fp *Foo) Panic() (err error) {
    defer PanicRecovery(&err)
    fp.Is[0] = 5
    return nil
}

func PanicRecovery(err *error) {

    if r := recover(); r != nil {
        if _, ok := r.(runtime.Error); ok {
                //fmt.Println("Panicing")
                //panic(r)
                *err = r.(error)
        } else {
            *err = r.(error)
        }
    }
}
```

The use of a separate function (rather than closure) allows re-use of the same function in other functions prone to panic.
