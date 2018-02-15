package main

import (
	"fmt"
)

// :show start
func worker(ch chan int, chQuit chan struct{}) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("Got value %d\n", v)
		case <-chQuit:
			fmt.Printf("Signalled on quit channel. Finishing\n")
			chQuit <- struct{}{}
			return
		}
	}
}
func main() {
	ch, chQuit := make(chan int), make(chan struct{})
	go worker(ch, chQuit)
	ch <- 3
	chQuit <- struct{}{}

	// wait to be signalled back by the worker
	<-chQuit
}

// :show end
