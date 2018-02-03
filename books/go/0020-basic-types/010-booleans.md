Title: Booleans
Search: bool
Id: rd6000m1
Body:

Type `bool` can be `true` or `false`.

```go
var b bool = true
fmt.Printf("b is: '%v'\n", b)
b = false
fmt.Printf("b is: '%v'\n", b)
```
prints:
```text
b is: 'true'
b is: 'false'
```

Size of `bool` variable (e.g. when part of a struct) is 1 byte.

```go
var b bool = true
fmt.Printf("size of bool is: %d\n", unsafe.Sizeof(b))
```
prints:
```text
size of bool is: 1
```

