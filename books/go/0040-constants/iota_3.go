package main

import "fmt"

func main() {
	// :show start
	const ( // iota is reset to 0
		a = 1 << iota // a == 1
		b = 1 << iota // b == 2
		c = 3         // c == 3  (iota is not used but still incremented)
		d = 1 << iota // d == 8
	)
	fmt.Printf("a: %d, b: %d, c: %d, d: %d\n", a, b, c, d)
	// :show end
}
