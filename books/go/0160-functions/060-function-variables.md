---
Title: Variables of function type
Id: 116
SOId: 801000g3
---

Functions are first-class values in Go:

* they can be assigned to variables
* they can be passed as values to functions

## Assigning functions to variables

@file function_variable.go output sha1:9be8e90aa90753a08b432b65421d14f76ac6b97d goplayground:2INQr00nmjX

Using variables for functions can be used when testing for mocking.

During real operation a variable points to a real function.

During tests a variable points to a function that mocks the functionality.

## Passing function as function arguments

@file function_variable2.go output sha1:51cf7bb8e0522a64b1065499f54add242dbe906a goplayground:Da9-TKksC1t

Common uses of passing functions:

* [filepath.Walk](https://golang.org/pkg/path/filepath/#Walk) takes a callback function to be called for every file found
* [ast.Inspect](https://golang.org/pkg/go/ast/#Inspect) traverses a tree and calls a function for each node

<!-- TODO: describe visitor pattern -->
