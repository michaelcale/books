package main

import "fmt"

// :show start

func logNum(i int) {
	fmt.Printf("Num %d\n", i)
}

func main() {
	i := 1
	defer logNum(i) // deferred function call: logNum(1)
	fmt.Println("First main statement")
	i++
	defer logNum(i)     // deferred function call: logNum(2)
	defer logNum(i * i) // deferred function call: logNum(4)
}

// :show end
