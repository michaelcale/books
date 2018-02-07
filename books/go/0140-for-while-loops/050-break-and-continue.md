---
Title: Break and Continue
Id: 3173
Score: 1
---
Breaking out of the loop and continuing to the next iteration is also supported in Go, like in many other languages:

```go
for x := 0; x < 10; x++ { // loop through 0 to 9
    if x < 3 { // skips all the numbers before 3
        continue
    }
    if x > 5 { // breaks out of the loop once x == 6
        break
    }
    fmt.Println("iteration", x)
}

// would print:
// iteration 3
// iteration 4
// iteration 5
```

The `break` and `continue` statements additionally accept an optional label, used to identify outer loops to target with the statement:

```go
OuterLoop:
for {
    for {
        if allDone() {
            break OuterLoop
        }
        if innerDone() {
            continue OuterLoop
        }
        // do something
    }
}
```
