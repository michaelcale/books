package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "is hello there?"
	toFind := "hello"
	if strings.Contains(s, toFind) {
		fmt.Printf("'%s' contains '%s'\n", s, toFind)
	} else {
		fmt.Printf("'%s' doesn't contain '%s'\n", s, toFind)
	}
	// :show end
}
