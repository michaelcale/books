package main

import "fmt"

func main() {
	// :show start
	var a int = 4
	pa := &a
	fmt.Printf("Address of a variable in memory is %p. Its value is: %d\n", pa, *pa)
	// :show end
}
