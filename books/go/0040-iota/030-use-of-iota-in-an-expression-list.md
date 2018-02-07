---
Title: Use of iota in an expression list
Id: 9707
Score: 4
---
Because `iota` is incremented after each [`ConstSpec`](https://golang.org/ref/spec#ConstSpec), values within the same expression list will have the same value for `iota`:

```go
const (
    bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0
    bit1, mask1                           // bit1 == 2, mask1 == 1
    _, _                                  // skips iota == 2
    bit3, mask3                           // bit3 == 8, mask3 == 7
)
```

This example was taken from the [Go Spec](https://golang.org/ref/spec#Iota) (CC-BY 3.0).
