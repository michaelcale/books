---
Title: Copy a map
Id: 9834
---
In Go all values are passed by copy.

However, similar to slices, the value of a map is only a reference to underlying data.

When you assign a map to another variable or pass to another function, you only copy the reference.

To copy the values, we need to write a bit of code:

@file copy_map.go
