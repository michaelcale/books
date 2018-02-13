package main

import "fmt"

// :show start
func main() {
	str := "Hello!"
	func() {
		fmt.Println(str)
	}()
}

// :show end
