---
Title: Testing Interface Implementation
Id: 9653
Score: 0
---
As Go uses implicit interface implementation, you will not get a compile-time error if your struct does not implement an interface you had intended to implement. You can test the implementation explicitly using type casting:

```go
    type MyInterface interface {
        Thing()
    }

    type MyImplementer struct {}

    func (m MyImplementer) Thing() {
        fmt.Println("Huzzah!")
    }

    // Interface is implemented, no error. Variable name _ causes value to be ignored.
    var _ MyInterface = (*MyImplementer)nil

    type MyNonImplementer struct {}

    // Compile-time error - cannot case because interface is not implemented.
    var _ MyInterface = (*MyNonImplementer)nil
```
