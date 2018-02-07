---
Title: Booleans
Search: bool
Id: rd6000m1
---
Type `bool` can be `true` or `false`.

```go
var b bool = true
fmt.Printf("b is: '%v'\n", b)
b = false
fmt.Printf("b is: '%v'\n", b)
```

**Output:**

```text
b is: 'true'
b is: 'false'
```

Size of `bool` variable (e.g. when part of a struct) is 1 byte.

```go
var b bool = true
fmt.Printf("size of bool is: %d\n", unsafe.Sizeof(b))
```

**Output:**

```text
size of bool is: 1
```

[Zero value](a-6069) of `bool` is `false`.
