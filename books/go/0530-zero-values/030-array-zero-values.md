Title: Array Zero Values
Id: 28500
Score: 1
Body:
As per the [Go blog][1]:
>  Arrays do not need to be initialized explicitly; the zero value of an array is a ready-to-use array whose elements are themselves zeroed 

For example, `myIntArray` is initialized with the zero value of `int`, which is 0:

    var myIntArray [5]int     // an array of five 0's: [0, 0, 0, 0, 0]


  [1]: https://blog.golang.org/go-slices-usage-and-internals
|======|
