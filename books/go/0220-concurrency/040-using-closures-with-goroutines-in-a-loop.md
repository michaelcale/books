---
Title: Using closures with goroutines in a loop
Id: 3897
Score: 5
---
When in a loop, the loop variable (val) in the following example is a single variable that changes value as it goes over the loop. Therefore one must do the following to actually pass each val of values to the goroutine:

```go
for val := range values {
    go func(val interface{}) {
        fmt.Println(val)
    }(val)
}
```

If you were to do just do go `func(val interface{}) { ... }()` without passing val, then the value of `val` will be whatever val is when the goroutines actually runs.

Another way to get the same effect is:

```go
for val := range values {
    val := val
    go func() {
        fmt.Println(val)
    }()
}
```

The strange-looking `val := val` creates a new variable in each iteration, which is then accessed by the goroutine.
