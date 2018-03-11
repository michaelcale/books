package main

import "fmt"

func fillAndCloseChannel(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i + 3
	}
	close(ch)
}

func main() {
	// :show start
	ch := make(chan int)
	go fillAndCloseChannel(ch)

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("v: %d\n", v)
	}
	// :show end
}
