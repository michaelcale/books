package main

import (
	"fmt"
	"log"
)

func main() {
	// :show start
	s := "1.2341"
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	if err != nil {
		log.Fatalf("fmt.Sscanf failed with '%s'\n", err)
	}
	fmt.Printf("f: %f\n", f)
	// :show end
}
