package main

import "fmt"

// :show start
func SumProduct(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	// only need the sum
	sum, _ := SumProduct(1, 2) // the product gets discarded
	fmt.Println(sum)           // prints 3
}

// :show end
