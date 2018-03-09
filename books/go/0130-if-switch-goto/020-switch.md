---
Title: switch statement
Id: 99
Score: 2
SOId: 4377
---
A simple `switch` statement:

```go
switch a + b {
case c:
    // do something
case d:
    // do something else
default:
    // do something entirely different
}
```

The above example is equivalent to:

```go
if a + b == c {
    // do something
} else if a + b == d {
    // do something else
} else {
    // do something entirely different
}
```

----------

The `default` clause is optional and will be executed if and only if none of the cases compare true, even if it does not appear last, which is acceptable.  The following is semantically the same as the first example:

```go
switch a + b {
default:
    // do something entirely different
case c:
    // do something
case d:
    // do something else
}
```

This could be useful if you intend to use the `fallthrough` statement in the `default` clause, which must be the last statement in a case and causes program execution to proceed to the next case:

```go
switch a + b {
default:
    // do something entirely different, but then also do something
    fallthrough
case c:
    // do something
case d:
    // do something else
}
```

----------

An empty switch expression is implicitly `true`:
```go
switch {
case a + b == c:
    // do something
case a + b == d:
    // do something else
}
```

----------

Switch statements support a simple statement similar to `if` statements:
```go
switch n := getNumber(); n {
case 1:
    // do something
case 2:
    // do something else
}
```

----------

Cases can be combined in a comma-separated list if they share the same logic:
```go
switch a + b {
case c, d:
    // do something
default:
    // do something entirely different
}
```

