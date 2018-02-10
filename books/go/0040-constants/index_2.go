package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s) // constant

	// A `const` statement can appear anywhere a `var` statement can.
	const n = 10
	fmt.Println(n)                           // 10
	fmt.Printf("n=%d is of type %T\n", n, n) // n=10 is of type int

	const m float64 = 4.3
	fmt.Println(m) // 4.3

	// An untyped constant takes the type needed by its context.
	// For example, here `math.Sin` expects a `float64`.
	const x = 10
	fmt.Println(math.Sin(x)) // -0.5440211108893699
}
