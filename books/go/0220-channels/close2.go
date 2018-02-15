package main

import (
	"fmt"
)

func main() {
	// :show start
	ch := make(chan string)
	close(ch)
	v := <-ch
	fmt.Printf("Receive from closed channel immediately returns zero value of the type: %#v\n", v)
	// :show end
}
