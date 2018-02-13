package main

import (
	"fmt"
)

// :show start
// addCheckOverflow adds two int16 numbers and additionally
// returns true if the result overflowed
func addCheckOverflow(a, b uint16) (uint16, bool) {
	res := a + b
	overflowed := res < a || res < b
	return res, overflowed
}

func main() {
	res, overflowed := addCheckOverflow(1, 3)
	fmt.Printf("%5d + %5d = %5d, overflowed: %v\n\n", 1, 3, res, overflowed)
	res, overflowed = addCheckOverflow(65520, 10000)
	fmt.Printf("%5d + %5d = %5d, overflowed: %v\n", 65550, 10000, res, overflowed)
}

// :show end
