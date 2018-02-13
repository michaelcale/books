package main

import "fmt"

// :show start
func gotoIfOdd(n int) {
	if n%2 == 1 {
		goto isOdd
	}
	fmt.Printf("%d is even\n", n)
	return

isOdd:
	fmt.Printf("%d is odd\n", n)
}

func main() {
	gotoIfOdd(5)
	gotoIfOdd(18)
}

// :show end
