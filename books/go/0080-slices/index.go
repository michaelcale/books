package main

import "fmt"

func main() {
	// :show start
	slice := make([]int, 0, 5)
	// append element to end of slice
	slice = append(slice, 5)
	// append multiple elements to end
	slice = append(slice, 3, 4)
	fmt.Printf("length of slice is: %d\n", len(slice))
	fmt.Printf("capacity of slice is: %d\n", cap(slice))
	// :show end
}
