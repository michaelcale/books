package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "original string original"
	s2 := strings.Replace(s, "original", "replaced", -1)
	fmt.Printf("s2: '%s'\n", s2)
	// :show end
}
