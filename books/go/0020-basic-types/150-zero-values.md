---
Title: Zero values
Id: 6069
---
Variables in Go are always initialized with a known value even if not explicitly assigned in source code.

Each Go type has a zero value.

Variables that are not explicitly initialized (assigned an explicit value) have value equal to zero value for their type.

This is different from C/C++, where variables that are not explicitly assigned have undefined values.

The values of zero type are unsurprising:

|type|zero value|
|----|----------|
|bool|false|
|integers|0|
|floating poing numbers|0.0|
|string|""|
|pointer|nil|
|slice|nil|
|map|nil|
|interface|nil|
|channel|nil|
|array|all elements have zero values|
|struct|all members set to zero value of their type|
|function|nil|

Said differently:
```go
var zeroBool bool                      // = false
var zeroInt int                        // = 0
var zeroF32 float32                    // = 0.0
var zeroStr string                     // = ""
var zeroPtr *int                       // = nil
var zeroSlice []uint32                 // = nil
var zeroMap map[string]int             // = nil
var zeroInterface interface{}          // = nil
var zeroChan chan bool                 // = nil
var zeroArray [5]int                   // = [0, 0, 0, 0, 0]
var zeroStruct struct{a int, b string} // = { a: 0, b: ""}
var zeroFunc func(bool)                // = nil
```
