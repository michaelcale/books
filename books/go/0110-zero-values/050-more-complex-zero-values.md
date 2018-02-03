Title: More Complex Zero Values
Id: 21170
Score: 0
Body:
In slices the zero value is an empty slice.

    var myIntSlice []int    // [] - an empty slice

Use `make` to create a slice populated with values, any values created in the slice are set to the zero value of the type of the slice. For instance:

    myIntSlice := make([]int, 5)    // [0, 0, 0, 0, 0] - a slice with 5 zeroes
    fmt.Println(myIntSlice[3])
    // Prints 0

In this example, `myIntSlice` is a `int` slice that contains 5 elements which are all 0 because that's the zero value for the type `int`.

You can also create a slice with `new`, this will create a pointer to a slice.

    myIntSlice := new([]int)        // &[] - a pointer to an empty slice
    *myIntSlice = make([]int, 5)    // [0, 0, 0, 0, 0] - a slice with 5 zeroes
    fmt.Println((*myIntSlice)[3])
    // Prints 0

_**Note:**_ Slice pointers don't support indexing so you can't access the values using `myIntSlice[3]`, instead you need to do it like `(*myIntSlice)[3]`.

|======|
