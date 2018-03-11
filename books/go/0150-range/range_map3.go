package main

import "fmt"

func main() {
	// :show start
	m := map[string]int{
		"three": 3,
		"five":  5,
	}
	for _, value := range m {
		fmt.Printf("value: %d\n", value)
	}
	// :show end
}
