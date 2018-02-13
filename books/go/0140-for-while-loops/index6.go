package main

import "fmt"

func main() {
	// :show start
	a := []int{1, 3, 5}
	for idx, value := range a {
		fmt.Printf("idx: %d, value: %d\n", idx, value)
	}
	// :show end
}
