Title: Simple use of iota
Id: 9704
Score: 6
Body:
To create a list of constants - assign `iota` value to each element:

    const (
      a = iota // a = 0
      b = iota // b = 1
      c = iota // c = 2
    )

To create a list of constants in a shortened way - assign `iota` value to the first element:

    const (
      a = iota // a = 0
      b        // b = 1
      c        // c = 2
    )
|======|
