package main

import (
	"fmt"
)

// :show start
func variadic(strs ...string) {
	// strs is a slice of string
	for i, str := range strs {
		fmt.Printf("%d: %s\n", i, str)
	}
	fmt.Print("\n")
}

func main() {
	variadic("Hello", "Goodbye")
	variadic("Str1", "Str2", "Str3")

	// you can also give a slice to a variadic function, with `...`:
	strs := []string{"a", "b"}
	variadic(strs...)
}

// :show end
