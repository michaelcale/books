Title: Basic Zero Values
Id: 21169
Body:
Variables in Go are always initialized whether you give them a starting value or not.  Each type, including custom types, has a zero value they are set to if not given a value.

    var myString string      // "" - an empty string
    var myInt int64          // 0 - applies to all types of int and uint
    var myFloat float64      // 0.0 - applies to all types of float and complex
    var myBool bool          // false
    var myPointer *string    // nil
    var myInter interface{}  // nil

This also applies to maps, slices, channels and function types.  These types will initialize to nil. In arrays, each element is initialized to the zero value of its respective type.
|======|
