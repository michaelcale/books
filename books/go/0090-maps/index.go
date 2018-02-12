package main

import "fmt"

func main() {
	// :show start
	// create an empty map
	m := make(map[string]int)

	// set the value
	m["three"] = 3
	m["four"] = 4

	// get the value and see if the value exists
	key := "four"
	if value, ok := m[key]; ok {
		fmt.Printf("Key '%s' exists and maps to %d\n", key, value)
	} else {
		fmt.Printf("Key '%s' doesn't exists\n", key)
	}

	key = "five"
	if value, ok := m[key]; ok {
		fmt.Printf("Key '%s' exists and maps to %d\n", key, value)
	} else {
		fmt.Printf("Key '%s' doesn't exists\n", key)
	}

	// if value doesn't exist, the result of lookup is zero value. In this case zero value of int is 0
	fmt.Printf("\nValue for non-existing key: %d\n\n", m["not-exists"])

	// iterating over keys and values
	fmt.Printf("All keys and their values:\n")
	for key, value := range m {
		fmt.Printf("%s => %d\n", key, value)
	}

	fmt.Printf("\nBefore deletion: len(m)=%d\n", len(m))
	delete(m, "four")
	fmt.Printf("After deletion: len(m)=%d\n", len(m))
	// :show end
}
