---
Title: Remove elements from slice
Id: 3420
---
If you need to remove one or more elements from a slice, or if you need to work with a sub-slice of an existing slice; you can use the following method.

> The following examples use a slice of type int, but will work for all types of slice.

So, we need a slice from which we will remove some elements:

    slice := []int{1, 2, 3, 4, 5, 6}
    // > [1 2 3 4 5 6]

We also need the indexes of the elements to remove:

    // index of first element to remove (corresponding to the '3' in the slice)
    var first = 2

    // index of last element to remove (corresponding to the '5' in the slice)
    var last = 4

And so we can "slice" the slice, removing undesired elements:

    // keeping elements from start to 'first element to remove' (not keeping first to remove),
    // removing elements from 'first element to remove' to 'last element to remove'
    // and keeping all others elements to the end of the slice
    newSlice1 := append(slice[:first], slice[last+1:]...)
    // > [1 2 6]

    // you can do using directly numbers instead of variables
    newSlice2 := append(slice[:2], slice[5:]...)
    // > [1 2 6]

    // Another way to do the same
    newSlice3 := slice[:first + copy(slice[first:], slice[last+1:])]
    // > [1 2 6]

    // same that newSlice3 with hard coded indexes (without use of variables)
    newSlice4 := slice[:2 + copy(slice[2:], slice[5:])]
    // > [1 2 6]

To remove only one element, just put the index of this element as the first AND as the last index to remove, like this:

    var indexToRemove = 3
    newSlice5 := append(slice[:indexToRemove], slice[indexToRemove+1:]...)
    // > [1 2 3 5 6]

    // hard-coded version:
    newSlice5 := append(slice[:3], slice[4:]...)
    // > [1 2 3 5 6]

And you can also remove elements from the beginning of the slice:

    newSlice6 := append(slice[:0], slice[last+1:]...)
    // > [6]

    // That can be simplified into
    newSlice6 := slice[last+1:]
    // > [6]

You can also remove some elements from the end of the slice:

    newSlice7 := append(slice[:first], slice[first+1:len(slice)-1]...)
    // > [1 2]

    // That can be simplified into
    newSlice7 := slice[:first]
    // > [1 2]

> If the new slice has to contain exactly the same elements as the first one, you can use the same thing but with `last := first-1`.
(This can be useful in cases where your indexes are previously computed)
