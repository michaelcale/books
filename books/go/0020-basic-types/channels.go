package main

import "fmt"

func main() {
	// :show start
	// create unbuffered channel of int values with capacity of 1
	ch := make(chan int)
	// start a new goroutine that sends value 3 over a channel
	go func() { ch <- 3 }()
	// read the value from a channel
	// it waits until goroutine above sends a value
	n := <-ch
	fmt.Printf("n: %d\n", n)
	// :show end
}
