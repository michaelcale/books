package main

import "fmt"

// :show start
func Add(a, b int) int {
	return a + b
}

func AddAndMultiply(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	sum, mult := AddAndMultiply(5, 8)
	fmt.Printf("5+8=%d, 5*8=%d\n", sum, mult)
	sum = Add(6, 12)
	fmt.Printf("6+12=%d\n", sum)
}

// :show end
