---
Title: Copy a map
Id: 74
SOId: 9834
---
In Go all values are passed by copy.

However, similar to slices, the value of a map is only a reference to underlying data.

When you assign a map to another variable or pass to another function, you only copy the reference.

To copy the values, we need to write a bit of code:

@file copy_map.go sha1:14f3c89f1b73df669e3ab1dfaeb6fc20f534cb97 goplayground:u_QPJ3gDicZ
