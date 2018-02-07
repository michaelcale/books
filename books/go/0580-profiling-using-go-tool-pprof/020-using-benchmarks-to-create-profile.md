---
Title: Using Benchmarks to Create Profile
Id: 25548
Score: 1
---
For a non-main packages as well as main, **instead of adding flags inside the code**, write **benchmarks** in the test package , for example:

```go
func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}
```

Then run the test with the profile flag

> go test -cpuprofile cpu.prof -bench=.

And the benchmarks will be run and create a prof file with filename cpu.prof (in the above example).
