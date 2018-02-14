package main

import "fmt"

// :show start

func main() {
	for i := 0; i < 2; i++ {
		defer func(i2 int) {
			fmt.Printf("%d\n", i2)
		}(i)
	}
}

// :show end
