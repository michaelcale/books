package main

import "fmt"

// :show start
func main() {
	// range of a string
	for idx, rune := range "str" {
		fmt.Printf("idx: %d,rune: %d\n", idx, rune)
	}
}

// :show end
