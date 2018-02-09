package main

import "fmt"

// :show start
func multipleReturn() (int, int) {
	return 1, 2
}

func multipleReturn2() (a int, b int) {
	a = 3
	b = 4
	return
}

func main() {
	x, y := multipleReturn()  // x = 1, y = 2
	w, z := multipleReturn2() // w = 3, z = 4
	fmt.Printf("x: %d, y: %d\nw: %d, z: %d\n", x, y, w, z)
	// :show end
}
