Title: Basic Type Conversion
Id: 9652
Score: 0
Body:
There are two basic styles of type conversion in Go:

    // Simple type conversion
    var x := Foo{}    // x is of type Foo
    var y := (Bar)Foo // y is of type Bar, unless Foo cannot be cast to Bar, then compile-time error occurs.
    // Extended type conversion
    var z,ok := x.(Bar)    // z is of type Bar, ok is of type bool - if conversion succeeded, z has the same value as x and ok is true. If it failed, z has the zero value of type Bar, and ok is false.
|======|
