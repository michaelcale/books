---
Title: Pointers
Id: 801000gd
---

A pointer to a type is address of the value of that type in memory.

Unlike C, Go doesn't have pointer arithmetic. You can take an address of a variable but you can't add or substract from a pointer.

Pointer basics:
```go
var a int = 4
var *pa = &a
fmt.Printf("Address of a variable in memory is 0x%p. It's value is: %d\n", pa, *pa)
```

[Zero value](a-6069) of a pointer is nil.

Learn more about [pointers](ch-1239).
