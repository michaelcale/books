package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// :show start
	b := true
	fmt.Printf("size of bool is: %d\n", unsafe.Sizeof(b))
	// :show end
}
