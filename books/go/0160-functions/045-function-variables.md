---
Title: Variables of function type
Id: 801000g3
---

Functions are first-class values in Go.

That means:
* they can be assigned to variables
* they can be passed as values to functions

## Assigning functions to variables

@file function_variable.go output

Using variables for functions can be used when testing for mocking.

During real operation a variable points to a real function.

During tests a variable points to a function that mocks the functionality.

## Passing function as function arguments

@file function_variable2.go output

Common uses of passing functions:
* [filepath.Walk](https://golang.org/pkg/path/filepath/#Walk) takes a callback function to be called for every file found
* [ast.Inspect](https://golang.org/pkg/go/ast/#Inspect) traverses a tree and calls a function for each node

<!-- TODO: describe visitor pattern -->