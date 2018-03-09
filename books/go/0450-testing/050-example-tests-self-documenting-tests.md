---
Title: Example tests (self documenting tests)
Id: 248
Score: 1
SOId: 4042
---
This type of tests make sure that your code compiles properly and will appear in the generated documentation for your project. In addition to that, the example tests can assert that your test produces proper output.

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

import "fmt"

func ExampleSum() {
    x := Sum(1, 2)
    fmt.Println(x)
    fmt.Println(Sum(-1, -1))
    fmt.Println(Sum(0, 0))

    // Output:
    // 3
    // -2
    // 0
}
```

To execute your test, run `go test` in the folder containing those files OR put those two files in a sub-folder named `sum` and then from the parent folder execute `go test ./sum`. In both cases you will get an output similar to this:

```text
ok      so/sum    0.005s
```

If you are wondering how this is testing your code, here is another example function, which actually fails the test:

```go
func ExampleSum_fail() {
    x := Sum(1, 2)
    fmt.Println(x)

    // Output:
    // 5
}
```

When you run `go test`, you get the following output:

```sh
$ go test
--- FAIL: ExampleSum_fail (0.00s)
got:
3
want:
5
FAIL
exit status 1
FAIL    so/sum    0.006s
```

If you want to see the documentation for your `sum` package – just run:

```sh
go doc -http=:6060
```

and navigate to http://localhost:6060/pkg/FOLDER/sum/, where _FOLDER_ is the folder containing the `sum` package (in this example `so`). The documentation for the sum method looks like this:

![enter image description here](http://i.stack.imgur.com/GNHv4.png)
