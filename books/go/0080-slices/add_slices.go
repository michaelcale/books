package main

import "fmt"

func main() {
	// :show start
	slice := []string{"!"}
	slice2 := []string{"Hello", "world"}
	slice = append(slice, slice2...)
	fmt.Printf("%#v\n", slice)
	// :show end
}
