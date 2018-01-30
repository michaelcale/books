Title: Skipping values
Id: 9706
Score: 3
Body:
The value of `iota` is still incremented for every entry in a constant list even if iota is not used:

    const ( // iota is reset to 0
        a = 1 << iota  // a == 1
        b = 1 << iota  // b == 2
        c = 3          // c == 3  (iota is not used but still incremented)
        d = 1 << iota  // d == 8
    )

it will also be incremented even if no constant is created at all, meaning the empty identifier can be used to skip values entirely:

    const (
      a = iota // a = 0
      _        // iota is incremented
      b        // b = 2
    )

The first code block was taken from the [Go Spec](https://golang.org/ref/spec#Iota) (CC-BY 3.0).


|======|
