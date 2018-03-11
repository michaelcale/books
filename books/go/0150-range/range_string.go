package main

import "fmt"

func main() {
	// :show start
	s := "Hey 世界"
	for idx, rune := range s {
		fmt.Printf("idx: %d, rune: %d\n", idx, rune)
	}
	// :show end
}
