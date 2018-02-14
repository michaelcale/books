package main

import "fmt"

// :show start
func logExit(name string) {
	fmt.Printf("Function %s returned\n", name)
}

func main() {
	fmt.Println("First main statement")
	defer logExit("main") // position of defer statement here does not matter
	fmt.Println("Last main statement")
}

// :show end
