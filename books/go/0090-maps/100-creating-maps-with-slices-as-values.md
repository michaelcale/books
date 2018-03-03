---
Title: Maps with slices as values
Id: 3910
---
    m := make(map[string][]int)

Accessing a non-existent key will return a nil slice as a value. Since nil slices act like zero length slices when used with `append` or other built-in functions you do not normally need to check to see if a key exists:

    // m["key1"] == nil && len(m["key1"]) == 0
    m["key1"] = append(m["key1"], 1)
    // len(m["key1"]) == 1

Deleting a key from a map sets the key back to a nil slice:

    delete(m, "key1")
    // m["key1"] == nil
