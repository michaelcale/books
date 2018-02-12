package main

import "fmt"

func main() {
	// :show start
	a := [3]int{4, 5} // array of 2 ints

	// access element of array
	fmt.Printf("a[2]: %d\n", a[2])

	// set element of array
	a[1] = 3

	// get size of array
	fmt.Printf("size of array a: %d\n", len(a))

	// when using [...] size will be deduced from { ... }
	a2 := [...]int{4, 8, -1} // array of 3 integers
	fmt.Printf("size of array a2: %d\n", len(a2))
	// :show end
}
