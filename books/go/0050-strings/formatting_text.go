package main

import (
	"fmt"
)

func main() {
	// :show start
	s := fmt.Sprintf("Hello %s", "World")
	fmt.Printf("s: '%s'\n", s)
	s = fmt.Sprintf("%d + %f = %d", 2, float64(3), 5)
	fmt.Println(s)
	// :show end
}
