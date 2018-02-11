package main

import "fmt"

func main() {
	// :show start
	var array = [6]int{1, 2, 3, 4, 5, 6}

	// doesn't work: invalid array index -1 (index must be non-negative)
	// fmt.Println(array[-42])
	fmt.Println(array[0]) // > 1
	fmt.Println(array[1]) // > 2
	fmt.Println(array[2]) // > 3
	fmt.Println(array[3]) // > 4
	fmt.Println(array[4]) // > 5
	fmt.Println(array[5]) // > 6
	// doesn't work: invalid array index 6 (out of bounds for 6-element array)
	//fmt.Println(array[6])
	// :show end
}
