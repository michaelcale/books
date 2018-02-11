package main

import "fmt"

func main() {
	// :show start
	slice := []int{1, 2, 3, 4}
	// create a zero-length slice with the same underlying array
	tmp := slice[:0]

	for _, v := range slice {
		if v%2 == 0 {
			// collect only wanted values
			tmp = append(tmp, v)
		}
	}
	fmt.Printf("%#v\n", tmp)
	// :show end
}
