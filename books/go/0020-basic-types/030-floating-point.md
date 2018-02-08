---
Title: Floating-point numbers
Search: float, double, float32, float64
Id: rd6000u7
---
Go has floating point numbers correspoinding to IEEE 754 standard:
* `float32` is 4 byte floating-point number (known as `float` in C)
* `float64` is 8 byte floating-point number (known as `double` in C)

[Zero value](a-6069) of `float32` and `float64` is 0.0.

## Converting floats to strings with `FormatFloat`

@file floating-point.go output

## Converting floats to strings with `Sprintf`

@file floating-point-2.go output

Using `strconv.FormatFloat` is faster than `fmt.Sprintf`.

## Converting string to float with `ParseFloat`

@file floating-point-3.go output

## Converting string to float with `Sscanf`

@file floating-point-4.go output
