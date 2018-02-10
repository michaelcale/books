package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "where hello is?"
	toFind := "hello"
	idx := strings.Index(s, toFind)
	fmt.Printf("'%s' is in s starting at position %d\n", toFind, idx)

	// when string is not found, result is -1
	idx = strings.Index(s, "not present")
	fmt.Printf("Index of non-existent substring is: %d\n", idx)
	// :show end
}
