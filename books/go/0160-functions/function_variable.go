package main

import "fmt"

// :show start
// intOp is a variable whose type is function that takes
// 2 integers as arguments and returns an integer
var intOp func(int, int) int

func intAdd(a, b int) int {
	return a + b
}

func main() {
	intOp = intAdd
	fmt.Printf("intOp(2, 3) = %d\n", intOp(2, 3))

	// we can assign literal functions as well
	intOp = func(a, b int) int {
		return a * b
	}
	fmt.Printf("intOp(2, 3) = %d\n", intOp(2, 3))
}

// :show end
