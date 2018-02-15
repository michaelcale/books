package main

import (
	"fmt"
	"time"
)

// :show start

func main() {
	chResult := make(chan int, 1)

	go func() {
		time.Sleep(1 * time.Second)
		chResult <- 5
		fmt.Printf("Worker finished")
	}()

	select {
	case res := <-chResult:
		fmt.Printf("Got %d from worker\n", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("Timed out before worker finished\n")
	}
}

// :show end
