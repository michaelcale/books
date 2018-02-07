---
Title: Zero value of slice
Id: 6807
Score: 1
---
The zero value of slice is `nil`, which has the length and capacity `0`. A `nil` slice has no underlying array. But there are also non-nil slices of length and capacity `0`, like `[]int{}` or `make([]int, 5)[5:]`.

Any type that have nil values can be converted to `nil` slice:

    s = []int(nil)

To test whether a slice is empty, use:

    if len(s) == 0　{
        fmt.Ptintf("s is empty.")
    }
