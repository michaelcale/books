package main

import "fmt"

func main() {
	// :show start
	a := []int{3, 15, 8}
	for _, el := range a {
		fmt.Printf("element: %d\n", el)
	}
	// :show end
}
