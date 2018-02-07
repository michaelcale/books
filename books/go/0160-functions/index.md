---
Title: Functions
Id: 373
---
## Introduction

Functions in Go provide organized, reusable code to perform a set of actions. Functions simplify the coding process, prevent redundant logic, and make code easier to follow. This topic describes the declaration and utilization of functions, arguments, parameters, return statements and scopes in Go.

## Syntax
- func() // function type with no arguments and no return value
- func(x int) int // accepts integer and returns an integer
- func(a, b int, z float32) bool // accepts 2 integers, one float and returns a boolean
- func(prefix string, values ...int) // "variadic" function which accepts one string and one or more number of integers
- func() (int, bool) // function returning two values
- func(a, b int, z float64, opt ...interface{}) (success bool) // accepts 2 integers, one float and one or more number of interfaces and returns named boolean value (which is already initialized inside of function)
