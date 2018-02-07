---
Title: Return Values
Id: 1252
Score: 1
---
A function can return one or more values to the caller:

```go
func AddAndMultiply(a, b int) (int, int) {
    return a+b, a*b
}
```

The second return value can also be the error var :

```go
import errors

func Divide(dividend, divisor int) (int, error) {
    if divisor == 0 {
        return 0, errors.New("Division by zero forbidden")
    }
    return dividend / divisor, nil
}
```

Two important things must be noted:

- The parenthesis may be omitted for a single return value.
- Each `return` statement must provide a value for **all** declared return values.

