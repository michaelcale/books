package main

import "fmt"

func main() {
	// :show start
	a := []int{3, 15, 8}
	for idx := range a {
		fmt.Printf("idx: %d\n", idx)
	}
	// :show end
}
