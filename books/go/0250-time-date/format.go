package main

import (
	"fmt"
	"time"
)

func main() {
	// :show start
	t := time.Date(2017, 9, 4, 3, 38, 45, 0, time.UTC)
	fmt.Println(t.Format("2006-02-01 15:04:05.000 MST"))
	fmt.Println(t.Format("2006-02-1 15pm"))
	fmt.Println(t.Format("Jan 06 Mon 2 01"))
	fmt.Println(t.Format("January 6 Mon 2 1"))
	fmt.Println(t.Format("Month: Jan '1', '01', _2"))
	// :show end
}
