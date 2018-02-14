package main

import (
	"fmt"
)

// :show start
// MyError is a custom error type
type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func printError(err error) {
	fmt.Printf("%s\n", err)
}

func main() {
	printError(&MyError{msg: "custom error type"})
}

// :show end
