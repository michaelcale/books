Title: Multiple constants declaration
Id: 9400
Score: 1
Body:
You can declare multiple constants within the same `const` block:

    const (
        Alpha = "alpha"
        Beta  = "beta"
        Gamma = "gamma"
    )

And automatically increment constants with the `iota` keyword:

    const (
        Zero = iota // Zero == 0
        One         // One  == 1
        Two         // Two  == 2
    )

For more examples of using `iota` to declare constants, see https://stackoverflow.com/documentation/go/2865/iota#t=20160724161043687765.

You can also declare multiple constants using the multiple assignment. However, this syntax may be harder to read and it is generally avoided.

    const Foo, Bar = "foo", "bar"
|======|
