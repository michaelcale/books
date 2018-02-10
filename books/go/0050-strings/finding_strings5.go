package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "this is string"
	toFind := "this"
	if strings.HasPrefix(s, toFind) {
		fmt.Printf("'%s' starts with '%s'\n", s, toFind)
	} else {
		fmt.Printf("'%s' doesn't start with '%s'\n", s, toFind)
	}
	// :show end
}
