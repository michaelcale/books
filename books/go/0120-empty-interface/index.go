package main

import (
	"fmt"
)

// :show start
func printVariableType(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("v is of type 'string'\n")
	case int:
		fmt.Printf("v is of type 'int'\n")
	default:
		// generic fallback
		fmt.Printf("v is of type '%T'\n", v)
	}
}

func main() {
	printVariableType("string") // string
	printVariableType(5)        // int
	printVariableType(int32(5)) // int32
}

// :show end
