---
Title: Constants
Search: const
Id: rd6000bc
---
You can declare constants of simple values that can be constructed at compile time.

Basic constant syntax:
```go
const (
    i int = 32   // int constant
    s = "string" // string constant
    i2 = 33      // untyped number constant

    // this, however, cannot be declared as a constant because []byte is
    // too complicated
    //m []byte = []byte{3, 4}
)
```

Learn more about [constants](ch-1047).
