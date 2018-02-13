package main

import (
	"fmt"
	"time"
)

func main() {
	// :show start
	for _ = range time.Tick(time.Second * 3) {
		fmt.Println("Ticking every 3 seconds")
	}
	// :show end
}
