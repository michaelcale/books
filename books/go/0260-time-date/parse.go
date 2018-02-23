package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// :show start
	s := "2017-04-09 03:38:45.000 UTC"
	t, err := time.Parse("2006-02-01 15:04:05.000 MST", s)
	if err != nil {
		log.Fatalf("time.Parse() failed wiht '%s'\n", err)
	}
	fmt.Printf("year: %d, month: %d, day: %d\n", t.Year(), t.Month(), t.Day())
	// :show end
}
