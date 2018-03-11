package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "Mixed Case"
	fmt.Printf("ToLower(s): '%s'\n", strings.ToLower(s))
	fmt.Printf("ToUpper(s): '%s'\n", strings.ToUpper(s))
	fmt.Printf("ToTitle(s): '%s'\n", strings.ToTitle(s))
	// :show end
}
