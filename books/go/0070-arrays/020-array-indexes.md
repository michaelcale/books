---
Title: Array indexes
Id: 32352
---
Array values should be accessed using a number specifying the location of the desired value in the array. This number is called the Index.

Indexes start at **0** and finish at **array length -1**.

To access a value, you have to do something like this: `arrayName[index]`, replacing "index" by the number corresponding to the offset of the value within your array.

For example:

@file array_indexes.go output sha1:33f5350a8cc56ec861d63ae233ab9295f6a9914d goplayground:6XD0m-9WltV


To set or modify a value in the array, you use the same index-based method.
Example:

@file array_indexes2.go output sha1:85b388de6ff2dcc3fbe68828164dbb4fdd96860d goplayground:ynkxqipMcV7
