package main

import (
	"fmt"
	"math"
)

// :show start
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("invalid argument '%f', must be >= 0", n)
	}
	return math.Sqrt(n), nil
}

func printSqrt(n float64) {
	if res, err := sqrt(n); err == nil {
		fmt.Printf("sqrt of %f is %f\n", n, res)
	} else {
		fmt.Printf("sqrt of %f returned error '%s'\n", n, err)
	}

}

func main() {
	printSqrt(16)
	printSqrt(-16)
}

// :show end
