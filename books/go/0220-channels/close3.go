package main

import (
	"fmt"
)

func main() {
	// :show start
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	v, isClosed := <-ch
	fmt.Printf("received %d, is channel closed: %v\n", v, isClosed)
	v, isClosed = <-ch
	fmt.Printf("received %d, is channel closed: %v\n", v, isClosed)
	// :show end
}
