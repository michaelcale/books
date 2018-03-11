package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "this is a string"
	a := strings.Split(s, " ")
	fmt.Printf("a: %#v\n", a)

	s2 := strings.Join(a, ",")
	fmt.Printf("s2: %#v\n", s2)
	// :show end
}
