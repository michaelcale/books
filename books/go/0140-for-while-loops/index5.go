package main

import "fmt"

func main() {
	// :show start
	i5 := 0
	for {
		fmt.Printf("i5: %d\n", i5)
		i5 += 2
		if i5 >= 5 {
			break
		}
	}

	// :show end
}
