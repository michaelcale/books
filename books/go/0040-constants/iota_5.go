package main

import "fmt"

func main() {
	// :show start
	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0
		bit1, mask1                          // bit1 == 2, mask1 == 1
		_, _                                 // skips iota == 2
		bit3, mask3                          // bit3 == 8, mask3 == 7
	)
	fmt.Printf("bit0: %d, mask0: 0x%x\n", bit0, mask0)
	fmt.Printf("bit3: %d, mask3: 0x%x\n", bit3, mask3)
	// :show end
}
