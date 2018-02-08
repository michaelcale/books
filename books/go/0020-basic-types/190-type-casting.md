---
Title: Type casting
Id: 80100098
---
Unlike a language like C, Go doesn't do implicit casting between types.

You have to cast between compatible types manually:

```go
// you can cast between numbers i.e. integers of various sizes and floating point
// numbers
var i1 int32 = 3
var i2 int = int(i1) // we must explicitly cast from int32 to int
var f float64 = float64(i1)

s := "string"
// we can cast between string and []byte and vice-versa
// note that unless optimizted by the compiler, this involves allocation
var d []byte = []byte(s)
```
