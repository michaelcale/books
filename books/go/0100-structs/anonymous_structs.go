package main

import "fmt"

func main() {
	// :show start
	data := struct {
		Number int
		Text   string
	}{
		42,
		"Hello world!",
	}

	fmt.Printf("data: %+v\n", data)
	// :show end
}
