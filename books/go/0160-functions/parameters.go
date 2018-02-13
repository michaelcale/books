package main

import (
	"fmt"
)

// :show start
func SayHelloToMe(firstName, lastName string, age int) {
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)
	fmt.Printf("You are %d", age)
}

func main() {
	SayHelloToMe("John", "Doe", 35)
}

// :show end
