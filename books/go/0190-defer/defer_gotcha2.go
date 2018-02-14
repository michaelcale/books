package main

import "fmt"

// :show start

func main() {
	for i := 0; i < 2; i++ {
		defer func() {
			fmt.Printf("%d\n", i)
		}()
	}
}

// :show end
