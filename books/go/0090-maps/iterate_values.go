package main

import "fmt"

func main() {
	// :show start
	people := map[string]int{
		"john": 30,
		"jane": 29,
		"mark": 11,
	}

	for _, value := range people {
		fmt.Printf("value: %d\n", value)
	}
	// :show end
}
