---
Title: Floating-point numbers
Search: float, double, float32, float64
Id: rd6000u7
---
Go has floating point numbers corresponding to the IEEE 754 standard:
* `float32` is a 4 byte floating-point number (known as `float` in C)
* `float64` is a 8 byte floating-point number (known as `double` in C)

[Zero value](a-6069) of `float32` and `float64` is 0.0.

## Convert floats to strings with `FormatFloat`

@file floating-point.go output sha1:56e9167e71cb198163749fa33494cf86cdf99072 goplayground:uql4Zc8Fiaz

## Convert floats to strings with `Sprintf`

@file floating-point-2.go output sha1:1051e611b391fe2419a211e2b54d0d6b5af137e6 goplayground:x25ikHfZM0T

Using `strconv.FormatFloat` is faster than `fmt.Sprintf`.

## Convert string to float with `ParseFloat`

@file floating-point-3.go output sha1:8d0c3f0446d9e36beab7c86e3f37420c336ea29f goplayground:EBIWcKWRw2p

## Convert string to float with `Sscanf`

@file floating-point-4.go output sha1:e9ddfc64925c7cbee478571a4b0ca3278c7bc421 goplayground:E3FzIjRMDmr
