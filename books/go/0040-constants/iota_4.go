package main

import "fmt"

func main() {
	// :show start
	const (
		a = iota // a = 0
		_        // iota is incremented
		b        // b = 2
	)
	fmt.Printf("a: %d, b: %d\n", a, b)
	// :show end
}
