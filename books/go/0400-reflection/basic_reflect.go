package main

import (
	"fmt"
	"reflect"
)

func main() {
	// :show start
	var v interface{} = 4
	var reflectVal reflect.Value = reflect.ValueOf(v)

	var typ reflect.Type = reflectVal.Type()
	fmt.Printf("Type '%s' of size: %d bytes\n", typ.Name(), typ.Size())
	if typ.Kind() == reflect.Int {
		fmt.Printf("v contains value of type int\n")
	}
	// :show end
}
