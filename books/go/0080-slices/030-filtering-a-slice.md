---
Title: Filtering a slice
Id: 6786
Score: 1
---
To filter a slice without allocating a new underlying array:

    // Our base slice
    slice := []int{ 1, 2, 3, 4 }
    // Create a zero-length slice with the same underlying array
    tmp := slice[:0]

    for _, v := range slice {
      if v % 2 == 0 {
        // Append desired values to slice
        tmp = append(tmp, v)
      }
    }

    // (Optional) Reassign the slice
    slice = tmp // [2, 4]
