package main

import (
	"fmt"
	"strconv"
)

func main() {
	// :show start
	s := "234"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("strconv.Atoi() failed with: '%s'\n", err)
	}
	fmt.Printf("strconv.Atoi('%s'): %d\n", s, i)

	i, err = strconv.Atoi("not a number")
	if err != nil {
		fmt.Printf("strconv.Atoi() failed with: '%s'\n", err)
	}

	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Printf("strconv.ParseInt() failed with: '%s'\n", err)
	}
	fmt.Printf("strconv.ParseInt('%s', 64): %d\n", s, i64)

	s = "-3.234"
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("strconv.ParseFloat() failed with: '%s'\n", err)
	}
	fmt.Printf("strconv.ParseFloat('%s', 64): %g\n", s, f64)

	var f2 float64
	_, err = fmt.Sscanf(s, "%f", &f2)
	if err != nil {
		fmt.Printf("fmt.Sscanf() failed with: '%s'\n", err)
	}
	fmt.Printf("fmt.Sscanf(): %g\n", f2)
	// :show end
}
