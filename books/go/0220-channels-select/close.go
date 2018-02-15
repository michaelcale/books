package main

import (
	"fmt"
	"time"
)

func main() {
	// :show start
	ch := make(chan string)

	go func() {
		for s := range ch {
			fmt.Printf("received from channel: %s\n", s)
		}
		fmt.Print("range loop finished because ch was closed\n")
	}()

	ch <- "foo"
	close(ch)
	// :show end

	// only to simplify example, don't sleep to coordinate
	// goroutines in real code
	time.Sleep(100 * time.Millisecond)
}
