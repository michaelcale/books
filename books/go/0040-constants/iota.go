package main

import "fmt"

func main() {
	// :show start
	const (
		Low = iota
		Medium
		High
	)
	fmt.Printf("Low: %d\nMedium: %d\nHigh: %d\n", Low, Medium, High)

	// :show end
}
