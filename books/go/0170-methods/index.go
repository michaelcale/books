package main

import "fmt"

// :show start
type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) PrintFullName() {
	fmt.Printf("%s %s", p.FirstName, p.LastName)
}

func main() {
	p := &Person{
		"John",
		"Doe",
	}
	p.PrintFullName()
}

// :show end
