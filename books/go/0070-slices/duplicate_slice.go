package main

import "fmt"

func main() {
	// :show start
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("src: %#v\n", src)
	fmt.Printf("dst: %#v\n", dst)
	// :show end
}
