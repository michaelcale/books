package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "hello and second hello"
	toFind := "hello"
	idx := strings.LastIndex(s, toFind)
	fmt.Printf("when searching from end, '%s' is in s at position %d\n", toFind, idx)
	// :show end
}
