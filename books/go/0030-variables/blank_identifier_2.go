package main

import "fmt"

// :show start
func main() {
	pets := []string{"dog", "cat", "fish"}

	// range returns both the current index and value
	// but sometimes we only need one or the other
	for _, pet := range pets {
		fmt.Println(pet)
	}
}

// :show end
