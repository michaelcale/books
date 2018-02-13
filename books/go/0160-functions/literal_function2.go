package main

import "fmt"

// :show start
func main() {
	func(str string) {
		fmt.Println(str)
	}("Hello!")
}

// :show end
