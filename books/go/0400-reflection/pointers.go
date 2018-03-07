package main

import (
	"fmt"
	"reflect"
)

// :show start
func printIntResolvingPointers(name string, v interface{}) {
	rv := reflect.ValueOf(v)
	typeName := rv.Type().String()
	for rv.Kind() == reflect.Ptr {
		name = "pointer to " + name
		rv = rv.Elem()
	}
	fmt.Printf("%s (%s) has value of %d\n", name, typeName, rv.Int())
}

func main() {
	n := 3
	printIntResolvingPointers("n", n)
	n = 4
	printIntResolvingPointers("n", &n)
	n = 5
	np := &n
	printIntResolvingPointers("n", &np)
}

// :show end
