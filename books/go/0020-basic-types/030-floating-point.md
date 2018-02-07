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

```go
var f32 float32 = 1.3
s1 := strconv.FormatFloat(float64(f32), 'E', -1, 32)
fmt.Printf("f32: %s\n", s1)
//
var f64 float64 = 8.1234
s2 := strconv.FormatFloat(f64, 'e', -1, 64)
fmt.Printf("f64: %s\n", s2)
```

**Output:**
```text
f32: 1.3E+00
f64: 8.1234e+00
```

## Converting floats to strings with `Sprintf`

```go
var f64 float64 = 1.54
s := fmt.Sprintf("%f", f64)
fmt.Printf("f is: '%s'\n", s)
```

**Output:**
```text
f is: '1.540000'
```

Using `strconv.FormatFloat` is faster than `fmt.Sprintf`.

## Converting string to float with `ParseFloat`

```go
s := "1.2341"
f, err := strconv.ParseFloat(s, 64)
if err != nil {
    log.Fatalf("strconv.ParseFloat() failed with '%s'\n", err)
}
fmt.Printf("f: %f\n", f)
```

**Output:**
```text
f: 1.234100
```

## Converting string to float with `Sscanf`

```go
s := "1.2341"
var f float64
_, err := fmt.Sscanf(s, "%f", &f)
if err != nil {
    log.Fatalf("fmt.Sscanf failed with '%s'\n", err)
}
fmt.Printf("f: %f\n", f)
```

**Output:**
```text
f: 1.234100
```
