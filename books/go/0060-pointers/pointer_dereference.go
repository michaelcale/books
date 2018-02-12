package main

import "fmt"

// :show start
type Person struct {
	Name string
}

func main() {
	c := new(Person) // returns pointer
	c.Name = "Catherine"
	fmt.Println(c.Name) // prints: Catherine
	d := c
	d.Name = "Daniel"
	fmt.Println(c.Name) // prints: Daniel
	// Adding an Asterix before a pointer dereferences the pointer
	i := *d
	i.Name = "Ines"
	fmt.Println(c.Name) // prints: Daniel
	fmt.Println(d.Name) // prints: Daniel
	fmt.Println(i.Name) // prints: Ines
	// :show end
}
