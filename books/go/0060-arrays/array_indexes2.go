package main

import "fmt"

func main() {
	// :show start
	var array = [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(array) // > [1 2 3 4 5 6]

	array[0] = 6
	fmt.Println(array) // > [6 2 3 4 5 6]

	array[1] = 5
	fmt.Println(array) // > [6 5 3 4 5 6]

	array[2] = 4
	fmt.Println(array) // > [6 5 4 4 5 6]

	array[3] = 3
	fmt.Println(array) // > [6 5 4 3 5 6]

	array[4] = 2
	fmt.Println(array) // > [6 5 4 3 2 6]

	array[5] = 1
	fmt.Println(array) // > [6 5 4 3 2 1]
	// :show end
}
