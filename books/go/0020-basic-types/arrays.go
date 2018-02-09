package main

import "fmt"

func main() {
	// :show start
	var a1 = [2]byte{3, 8} // array of 2 bytes
	// when using [...] size will be deduced from { ... }
	a2 := [...]int{1, 2, 3} // array of 3 integers

	fmt.Printf("Size of a1: %d.\nSize of a2: %d\n", len(a1), len(a2))
	// :show end
}
