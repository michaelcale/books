---
Title: Basic Test
Id: 4039
Score: 7
---

`main.go`:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(Sum(4,5))
}

func Sum(a, b int) int {
    return a + b
}
```

`main_test.go`:

```go
package main

import (
    "testing"
)

// Test methods start with `Test`
func TestSum(t *testing.T) {
    got := Sum(1, 2)
    want := 3
    if got != want {
        t.Errorf("Sum(1, 2) == %d, want %d", got, want)
    }
}
```

To run the test just use the `go test` command:

```sh
$ go test
ok      test_app    0.005s
```

Use the `-v` flag to see the results of each test:

```sh
$ go test -v
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
ok      _/tmp    0.000s
```

Use the path `./...` to test subdirectories recursively:

```sh
$ go test -v ./...
ok      github.com/me/project/dir1    0.008s
=== RUN   TestSum
--- PASS: TestSum (0.00s)
PASS
ok      github.com/me/project/dir2    0.008s
=== RUN   TestDiff
--- PASS: TestDiff (0.00s)
PASS
```

**Run a Particular Test:**
If there are multiple tests and you want to run a specific test, it can be done like this:

```sh
go test -v -run=<TestName> // will execute only test with this name
```

Example:

```sh
go test -v run=TestSum
```