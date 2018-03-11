package main

import "fmt"

func main() {
	// :show start
	a := []int{3, 15, 8}
	for idx, el := range a {
		fmt.Printf("idx: %d, element: %d\n", idx, el)
	}
	// :show end
}
