---
Title: Creating goroutines
Id: 1258
Score: 7
---
Any function can be invoked as a goroutine by prefixing its invocation with the keyword `go`:

```go
func DoMultiply(x,y int) {
    // Simulate some hard work
    time.Sleep(time.Second * 1)
    fmt.Printf("Result: %d\n", x * y)
}

go DoMultiply(1,2) // first execution, non-blocking
go DoMultiply(3,4) // second execution, also non-blocking

// Results are printed after a single second only,
// not 2 seconds because they execute concurrently:
// Result: 2
// Result: 12
```

Note that the return value of the function is ignored.

