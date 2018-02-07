---
Title: If Statements
Id: 4378
Score: 2
---
A simple `if` statement:

```go
if a == b {
    // do something
}
```

Note that there are no parentheses surrounding the condition and that the opening curly brace `{` must be on the same line.  The following will *not* compile:

```go
if a == b
{
    // do something
}
```

----------

An `if` statement making use of `else`:

```go
if a == b {
    // do something
} else if a == c {
    // do something else
} else {
    // do something entirely different
}
```

----------

Per [golang.org's documentation](https://golang.org/ref/spec#If_statements), "The expression may be preceded by a simple statement, which executes before the expression is evaluated."  Variables declared in this simple statement are scoped to the `if` statement and cannot be accessed outside it:

```go
if err := attemptSomething(); err != nil {
    // attemptSomething() was successful!
} else {
    // attemptSomething() returned an error; handle it
}
fmt.Println(err) // compiler error, 'undefined: err'
```
