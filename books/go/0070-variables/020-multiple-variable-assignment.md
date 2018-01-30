Title: Multiple Variable Assignment
Id: 2193
Score: 5
Body:
In Go, you can declare multiple variables at the same time.

    // You can declare multiple variables of the same type in one line
    var a, b, c string

    var d, e string = "Hello", "world!"

    // You can also use short declaration to assign multiple variables
    x, y, z := 1, 2, 3

    foo, bar := 4, "stack" // `foo` is type `int`, `bar` is type `string`

If a function returns multiple values, you can also assign values to variables based on the function's return values.

    func multipleReturn() (int, int) {
        return 1, 2
    }

    x, y := multipleReturn() // x = 1, y = 2

    func multipleReturn2() (a int, b int) {
        a = 3
        b = 4
        return
    }

    w, z := multipleReturn2() // w = 3, z = 4
|======|
