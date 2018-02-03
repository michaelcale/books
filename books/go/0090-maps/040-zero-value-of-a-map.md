Title: Zero value of a map
Id: 2485
Score: 9
Body:
The zero value of a `map` is `nil` and has a length of `0`.

    var m map[string]string
    fmt.Println(m == nil) // true
    fmt.Println(len(m) ==0) // true

A `nil` map has no keys nor can keys be added. A `nil` map behaves like an empty map if read from but causes a runtime panic if written to.

    var m map[string]string
    
    // reading
    m["foo"] == "" // true. Remember "" is the zero value for a string
    _, ok = m["foo"] // ok == false
    
    // writing
    m["foo"] = "bar" // panic: assignment to entry in nil map

You should not try to read from or write to a zero value map. Instead, initialize the map (with `make` or assignment) before using it.

    var m map[string]string
    m = make(map[string]string) // OR m = map[string]string{}
    m["foo"] = "bar"
|======|
