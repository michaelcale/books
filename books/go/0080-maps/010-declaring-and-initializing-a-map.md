Title: Declaring and initializing a map
Id: 2483
Score: 30
Body:
You define a map using the keyword `map`, followed by the types of its keys and its values:

    // Keys are ints, values are ints.
    var m1 map[int]int // initialized to nil

    // Keys are strings, values are ints.
    var m2 map[string]int // initialized to nil

Maps are reference types, and once defined they have a [_zero value_ of `nil`](http://stackoverflow.com/documentation/go/732/maps/2485/zero-value-of-a-map#t=201607220859253128466). Writes to nil maps will [panic](http://stackoverflow.com/documentation/go/4350/panic-and-recover/17009/panic#t=201608011756369491437) and reads will always return the zero value.

To initialize a map, use the [`make`](https://golang.org/pkg/builtin/#make) function:

    m := make(map[string]int)

With the two-parameter form of `make`, it's possible to specify an initial entry capacity for the map, overriding the default capacity:

    m := make(map[string]int, 30)

Alternatively, you can declare a map, initializing it to its zero value, and then assign a literal value to it later, which helps if you marshal the struct into json thereby producing an empty map on return.

    m := make(map[string]int, 0)

You can also make a map and set its initial value with curly brackets (`{}`).

    var m map[string]int = map[string]int{"Foo": 20, "Bar": 30}

    fmt.Println(m["Foo"]) // outputs 20

All the following statements result in the variable being bound to the same value.

    // Declare, initializing to zero value, then assign a literal value.
    var m map[string]int
    m = map[string]int{}

    // Declare and initialize via literal value.
    var m = map[string]int{}

    // Declare via short variable declaration and initialize with a literal value.
    m := map[string]int{}

We can also use a _map literal_ to [create a new map with some initial key/value pairs](http://stackoverflow.com/documentation/go/732/maps/2484/creating-a-map#t=201607220900510314955).

The key type can be any [comparable](http://golang.org/ref/spec#Comparison_operators) type; notably, [this excludes functions, maps, and slices](https://golang.org/ref/spec#Map_types). The value type can be any type, including custom types or `interface{}`.

    type Person struct {
        FirstName string
        LastName  string
    }

    // Declare via short variable declaration and initialize with make.
    m := make(map[string]Person)

    // Declare, initializing to zero value, then assign a literal value.
    var m map[string]Person
    m = map[string]Person{}

    // Declare and initialize via literal value.
    var m = map[string]Person{}

    // Declare via short variable declaration and initialize with a literal value.
    m := map[string]Person{}


|======|
