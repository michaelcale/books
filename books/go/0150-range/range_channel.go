package main

import "fmt"

// :show start
func fillAndCloseChannel(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i + 3
	}
	close(ch)
}

// :show end

func main() {
	// :show start
	ch := make(chan int)
	go fillAndCloseChannel(ch)

	for v := range ch {
		fmt.Printf("v: %d\n", v)
	}
	// :show end
}
