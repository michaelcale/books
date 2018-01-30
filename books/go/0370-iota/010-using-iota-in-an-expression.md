Title: Using iota in an expression
Id: 9705
Score: 14
Body:
`iota` can be used in expressions, so it can also be used to assign values other than simple incrementing integers starting from zero. To create constants for SI units, use this example from [Effective Go][1]:

    type ByteSize float64

    const (
        _           = iota // ignore first value by assigning to blank identifier
        KB ByteSize = 1 << (10 * iota)
        MB
        GB
        TB
        PB
        EB
        ZB
        YB
    )


  [1]: https://golang.org/doc/effective_go.html#initialization
|======|
