package main

import "fmt"

// :show start

func plusOne(i int) (result int) {
	// anonymous function must be called by adding ()
	defer func() { result++ }()

	// i is returned as result, which is updated by deferred function above
	// after execution of below return
	return i
}

func main() {
	fmt.Println(plusOne(1)) // 2
}

// :show end
