package main

import "fmt"

// :show start
func inverse(v float32) (reciprocal float32) {
	if v == 0 {
		return
	}
	reciprocal = 1 / v
	return
}

// function can return multiple values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Printf("inverse(5)=%.2f\n", inverse(5))
}

// :show end
