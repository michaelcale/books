Title: Blank Identifier
Id: 29103
Score: 2
Body:
Go will throw an error when there is a variable that is unused, in order to encourage you to write better code. However, there are some situations when you really don't need to use a value stored in a variable. In those cases, you use a "blank identifier" `_` to assign and discard the assigned value.

A blank identifier can be assigned a value of any type, and is most commonly used in functions that return multiple values.

**Multiple Return Values**

```go
func SumProduct(a, b int) (int, int) {
    return a+b, a*b
}

func main() {
    // I only want the sum, but not the product
    sum, _ := SumProduct(1,2) // the product gets discarded
    fmt.Println(sum) // prints 3
}
```

**Using `range`**

```go
func main() {

    pets := []string{"dog", "cat", "fish"}

    // Range returns both the current index and value
    // but sometimes you may only want to use the value
    for _, pet := range pets {
        fmt.Println(pet)
    }

}
```
|======|
