package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// :show start
	s := "1.2341"
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("strconv.ParseFloat() failed with '%s'\n", err)
	}
	fmt.Printf("f: %f\n", f)
	// :show end
}
