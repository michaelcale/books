package main

import "fmt"

func main() {
	// :show start
	i3 := 0
	for ; ; i3 += 2 {
		fmt.Printf("i3: %d\n", i3)
		if i3 >= 5 {
			break
		}
	}
	// :show end
}
