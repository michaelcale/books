package main

import "fmt"

// :show start
func retVals(ok bool) (int, bool) {
	return 5, ok
}

func main() {
	a := 5
	b := 5
	if a == b {
		fmt.Print("a is equal to b\n")
	} else {
		fmt.Print("a is not equal to b\n")
	}

	if n, ok := retVals(true); ok {
		fmt.Print("ok is true, n: %d\n", n)
	} else {
		fmt.Print("ok is false, n: %d\n", n)
	}

	// n is only visible within if loop, so this would fail compilation
	// fmt.Printf("n: %d\n", n)
}

// :show end
