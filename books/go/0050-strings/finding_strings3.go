package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "first is, second is, third is"
	toFind := "is"
	currStart := 0
	for {
		idx := strings.Index(s, toFind)
		if idx == -1 {
			break
		}
		fmt.Printf("found '%s' at position %d\n", toFind, currStart+idx)
		currStart += idx + len(toFind)
		s = s[idx+len(toFind):]
	}
	// :show end
}
