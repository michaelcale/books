package main

import "fmt"

func main() {
	// :show start
	var s string // empty string ""
	s1 := "string\nliteral\nwith\tescape characters"
	s2 := `raw string literal
	which doesn't recgonize escape characters like \n
	`

	// you can add strings with +
	fmt.Printf("sum of string: %s\n", s+s1+s2)

	// you can compare strings with ==
	if s1 == s2 {
		fmt.Printf("s1 is equal to s2\n")
	} else {
		fmt.Printf("s1 is not equal to s2\n")
	}

	fmt.Printf("substring of s1: %s\n", s1[3:5])
	fmt.Printf("byte (character) at position 3 in s1: %d\n", s1[3])

	// C-style string formatting
	s = fmt.Sprintf("%d + %f = %s", 1, float64(3), "4")
	fmt.Printf("s: %s\n", s)
	// :show end
}
