package main

import (
	"fmt"
	"reflect"
)

// :show start
func printIntResolvingPointers(v interface{}) {
	rv := reflect.ValueOf(v)
	typeName := rv.Type().String()
	name := ""
	for rv.Kind() == reflect.Ptr {
		name = "pointer to " + name
		rv = rv.Elem()
	}
	name += rv.Type().String()
	fmt.Printf("Value: %d. Type: '%s' i.e. '%s'.\n\n", rv.Int(), name, typeName)
}

func main() {
	n := 3
	printIntResolvingPointers(n)

	n = 4
	printIntResolvingPointers(&n)

	n = 5
	np := &n
	printIntResolvingPointers(&np)
}

// :show end
