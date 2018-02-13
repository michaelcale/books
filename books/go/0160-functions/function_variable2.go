package main

import "fmt"

// :show start
func funcAdd(a, b int) int {
	return a + b
}

func runFunc(a, b int, intOp func(int, int) int) {
	fmt.Printf("intOp(%d, %d) = %d\n", a, b, intOp(a, b))
}

func main() {
	runFunc(2, 3, funcAdd)

	// we can pass literal functions as well
	runFunc(2, 3, func(a, b int) int {
		return a * b
	})
}

// :show end
