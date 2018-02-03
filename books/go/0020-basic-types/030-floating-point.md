Title: Floating-point numbers
Search: float, double, float32, float64
Id: rd6000u7
Body:
Go has floating point numbers correspoinding to IEEE 754 standard:
* `float32` is 4 byte floating-point number (known as `float` in C)
* `float64` is 8 byte floating-point number (known as `double` in C)

[Zero value](a-6069) of `float32` and `float64` is 0.0.

## Converting floats to strings

```go
var f64 float64 = 1.54
s := fmt.Sprintf("%f", f64)
fmt.Printf("f is: '%s'\n", s)
```

## Converting string to float

TODO: write me
