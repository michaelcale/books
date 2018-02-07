---
Title: Copying contents from one slice to another slice
Id: 3749
Score: 0
---
If you wish to copy the contents of a slice into an initially empty slice, following steps can be taken to accomplish it-

1) Create the source slice:


    var sourceSlice []interface{} = []interface{}{"Hello",5.10,"World",true}

2) Create the destination slice, with:

- Length = Length of sourceSlice


    var destinationSlice []interface{} = make([]interface{},len(sourceSlice))

3) Now that the destination slice's underlying array is big enough to accomodate all the elements of the source slice, we can proceed to copy the elements using the builtin `copy`:

    copy(destinationSlice,sourceSlice)
