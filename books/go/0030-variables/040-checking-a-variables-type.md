---
Title: Checking a variable's type
Id: 29105
---
There are some situations where you won't be sure what type a variable is when it is returned from a function. You can always check a variable's type by using `var.(type)` if you are unsure what type it is:

```go
x := someFunction() // Some value of an unknown type is stored in x now

switch x := x.(type) {
    case bool:
        fmt.Printf("boolean %t\n", x)             // x has type bool
    case int:
        fmt.Printf("integer %d\n", x)             // x has type int
    case string:
        fmt.Printf("pointer to boolean %s\n", x) // x has type string
    default:
        fmt.Printf("unexpected type %T\n", x)     // %T prints whatever type x is
}
```
