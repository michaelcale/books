package main

import (
	"fmt"
	"time"
)

// :show start
func mult(x, y int) {
	fmt.Printf("%d * %d = %d\n", x, y, x*y)
}

// :show end

func main() {
	// :show start
	go mult(1, 2) // first execution, non-blocking
	go mult(3, 4) // second execution, also non-blocking
	// :show end

	// that's not how you do it in real code
	time.Sleep(200 * time.Millisecond)
}
