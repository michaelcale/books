---
Title: Variables
Id: 674
---
Variable declarations follows **variable-name** **variable-type** format.

It's a different order than C. On the plus side it's more consistent and complex delcarations are easier to write.

Various ways of defining variables:
```go
// declaration of a single top-level variable
var topLevelInt int = 5

// grouping of multiple top-level declarations
var (
    intVal int // value is initialized with zero-value
    str string = "str" // assigning

    // functions are first-class values so can be assigned to variables
    // f is variable of type func(a int) string
    // it's uninitialized so is nil (zero-value for function variables)
    f func(a int) string

)

func f() {
    // shorthand using local type inferenc
    // type of `i` is int and is infered from the value
    // note: this is not allowed at top-level
    i := 4

    // grouping inside a function
    var (
        i int
        s string
    )

    // _ is like a variable whose value is discarded. It's called blank identifier.
    // Useful when we don't care about one of the values returned by a function
    _, err := io.Copy(dst, src) // don't care how many bytes were written
    ...
}
```
