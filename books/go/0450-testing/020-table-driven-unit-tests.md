---
Title: Table-driven unit tests
Id: 4041
Score: 6
---
This type of testing is popular technique for testing with predefined input and output values.

Create a file called `main.go` with content:

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(Sum(4, 5))
}

func Sum(a, b int) int {
    return a + b
}
```

After you run it with, you will see that the output is `9`. Although the `Sum` function looks pretty simple, it is a good idea to test your code. In order to do this, we create another file named `main_test.go` in the same folder as `main.go`, containing the following code:

```go
package main

import (
    "testing"
)

// Test methods start with Test
func TestSum(t *testing.T) {
    // Note that the data variable is of type array of anonymous struct,
    // which is very handy for writing table-driven unit tests.
    data := []struct {
        a, b, res int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {1, -1, 0},
        {2, 3, 5},
        {1000, 234, 1234},
    }

    for _, d := range data {
        if got := Sum(d.a, d.b); got != d.res {
            t.Errorf("Sum(%d, %d) == %d, want %d", d.a, d.b, got, d.res)
        }
    }
}
```

As you can see, a slice of anonymous structs is created, each with a set of inputs and the expected result. This allows a large number of test cases to be created all together in one place, then executed in a loop, reducing code repetition and improving clarity.
