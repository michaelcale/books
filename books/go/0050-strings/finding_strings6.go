package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "this is string"
	toFind := "string"
	if strings.HasSuffix(s, toFind) {
		fmt.Printf("'%s' ends with '%s'\n", s, toFind)
	} else {
		fmt.Printf("'%s' doesn't end with '%s'\n", s, toFind)
	}
	// :show end
}
