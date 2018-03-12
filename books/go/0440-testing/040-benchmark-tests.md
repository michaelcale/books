---
Title: Benchmark tests
Id: 247
Score: 2
SOId: 4040
---
If you want to measure benchmarks add a testing method like this:

`sum.go`:
```go
package sum

// Sum calculates the sum of two integers
func Sum(a, b int) int {
    return a + b
}
```

`sum_test.go`:

```go
package sum

import "testing"

func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = Sum(2, 3)
    }
}
```

Then in order to run a simple benchmark:

```sh
$ go test -bench=.
BenchmarkSum-8    2000000000             0.49 ns/op
ok      so/sum    1.027s
```
