package main

import "fmt"

func main() {
	// :show start
	m := map[string]int{
		"three": 3,
		"five":  5,
	}
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	// :show end
}
