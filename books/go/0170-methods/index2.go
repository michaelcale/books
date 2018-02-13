package main

import "fmt"

// :show start
type Person struct {
	FirstName string
	LastName  string
}

func (p Person) PrintFullNameValue() {
	fmt.Printf("PrintFullNameValue:   address of p is %p\n", &p)
}

func (p *Person) PrintFullNamePointer() {
	fmt.Printf("PrintFullNamePointer: p is            %p\n", p)
}

func main() {
	p := Person{
		"John",
		"Doe",
	}
	fmt.Printf("address of p:                         %p\n", &p)
	p.PrintFullNamePointer()
	p.PrintFullNameValue()
}

// :show end
