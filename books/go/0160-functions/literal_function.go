package main

import "fmt"

// :show start
func main() {
	func() {
		fmt.Println("Hello!")
	}()
}

// :show end
