package main

import (
	"fmt"
	"os"
)

// :show start
func main() {
	path := "file_not_exists.go"
	err := os.Remove(path)
	if err != nil {
		fmt.Printf("os.Remove failed with '%s'\n", err)
		fmt.Printf("is not exists: %v\n", os.IsNotExist(err))
	}
}

// :show end
