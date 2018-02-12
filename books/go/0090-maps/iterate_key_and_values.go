package main

import "fmt"

func main() {
	// :show start
	people := map[string]int{
		"john": 30,
		"jane": 29,
		"mark": 11,
	}

	for key, value := range people {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	// :show end
}
