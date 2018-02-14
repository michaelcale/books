package main

import "fmt"

// :show start

func main() {
	fmt.Print("Before if\n")
	if true {
		defer fmt.Print("inside if\n")
	}

	fmt.Print("Ater if\n")
}

// :show end
