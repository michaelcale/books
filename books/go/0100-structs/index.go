package main

import (
	"fmt"
)

// :show start
// Person describes a person
type Person struct {
	FirstName string
	LastName  string
}

// FullName returns full name of a person
func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func main() {
	// zero value of struct
	var p Person
	fmt.Printf("p: %v\n\n", p)

	p = Person{
		FirstName: "John",
		LastName:  "Doe",
	}
	fmt.Printf("p: %v\n\n", p)

	// call a method on a struct
	fmt.Printf("p.FullName(): %s\n", p.FullName())
}

// :show end
