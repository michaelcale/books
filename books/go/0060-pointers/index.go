package main

import "fmt"

func main() {
	// :show start
	v := 5

	// pv is a pointer to v
	pv := &v
	fmt.Printf("v: %d, pv: %p\n", v, pv)

	// we change the value of v via pv
	*pv = 4
	fmt.Printf("v: %d\n", v)

	// two pointers to the same value have the same address
	pv2 := &v
	fmt.Printf("pv: %p, pv2: %p\n", pv, pv2)
	// :show end
}
