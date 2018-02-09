package main

import "fmt"

func main() {
	// :show start
	s := "日本語"
	for i, runeChar := range s {
		fmt.Printf("Rune at byte position %d is %#U\n", i, runeChar)
	}
	// :show end
}
