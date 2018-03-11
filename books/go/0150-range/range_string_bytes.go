package main

import "fmt"

func main() {
	// :show start
	s := "Hey 世界"
	for idx := range s {
		b := s[idx]
		fmt.Printf("idx: %d, byte: %d\n", idx, b)
	}
	// :show end
}
