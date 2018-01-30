Title: Basic Loop
Id: 3172
Score: 7
Body:
`for` is the only loop statement in go, so a basic loop implementation could look like this:

    // like if, for doesn't use parens either.
    // variables declared in for and if are local to their scope.
    for x := 0; x < 3; x++ { // ++ is a statement.
        fmt.Println("iteration", x)
    }

    // would print:
    // iteration 0
    // iteration 1
    // iteration 2
|======|
