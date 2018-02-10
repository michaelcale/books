package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s1 := "gone"
	s2 := "GoNe"
	if strings.EqualFold(s1, s2) {
		fmt.Printf("'%s' is equal '%s' when ignoring case\n", s1, s2)
	} else {
		fmt.Printf("'%s' is not equal '%s' when ignoring case\n", s1, s2)
	}
	// :show end
}
