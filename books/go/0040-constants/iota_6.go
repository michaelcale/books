package main

import "fmt"

func main() {
	// :show start
	type ByteSize int

	const (
		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)
	fmt.Printf("KB: %d\n", KB)
	// :show end
}
