package main

import (
	"fmt"
	"strings"
)

func main() {
	// :show start
	s := "  str  "
	fmt.Printf("TrimSpace('%s'): '%s'\n\n", s, strings.TrimSpace(s))

	s = "abacdda"
	cutset := "ab"
	fmt.Printf("Trim('%s', '%s'): '%s'\n\n", s, cutset, strings.Trim(s, cutset))

	fmt.Printf("TrimLeft('%s', '%s'): '%s'\n\n", s, cutset, strings.TrimLeft(s, cutset))

	fmt.Printf("TrimRight('%s', '%s'): '%s'\n\n", s, cutset, strings.TrimRight(s, cutset))

	suffix := "ab"
	fmt.Printf("TrimSuffix('%s', '%s'): '%s'\n\n", s, suffix, strings.TrimSuffix(s, suffix))

	prefix := "ab"
	fmt.Printf("TrimPrefix('%s', '%s'): '%s'\n\n", s, prefix, strings.TrimPrefix(s, prefix))
	// :show end
}
