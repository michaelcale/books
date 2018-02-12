package main

import "fmt"

func main() {
	// :show start
	greetings := map[string]struct{}{
		"hi":    {},
		"hello": {},
	}

	// delete a value from set
	delete(greetings, "hi")

	// add a value to set
	greetings["hey"] = struct{}{}

	// check if a value is in the set:
	if _, ok := greetings["hey"]; ok {
		fmt.Printf("'hey' is in greetings\n")
	}
	// :show end
}
