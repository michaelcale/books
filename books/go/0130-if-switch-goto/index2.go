package main

import "fmt"

func main() {
	// :show start
	stmt := "if"
	switch stmt {
	case "if", "for":
		fmt.Printf("stmt ('%s') is either 'if' or 'for'\n", stmt)
	case "else":
		fmt.Printf("stmt is 'else'\n")
	default:
		fmt.Printf("stmt is '%s'\n", stmt)
	}
	// :show end
}
