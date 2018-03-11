package main

import "fmt"

func main() {
	// :show start
	m := map[string]int{
		"three": 3,
		"five":  5,
	}
	for key := range m {
		fmt.Printf("key: %s\n", key)
	}
	// :show end
}
