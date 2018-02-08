package main

import "fmt"

func main() {
	// :show start
	var f64 float64 = 1.54
	s := fmt.Sprintf("%f", f64)
	fmt.Printf("f is: '%s'\n", s)
	// :show end
}
