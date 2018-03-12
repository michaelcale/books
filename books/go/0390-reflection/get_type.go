package main

import (
	"fmt"
	"reflect"
)

// :show start

func printType(v interface{}) {
	rv := reflect.ValueOf(v)
	typ := rv.Type()
	typeName := ""
	switch rv.Kind() {
	case reflect.Ptr:
		typeName = "pointer"
	case reflect.Int:
		typeName = "int"
	case reflect.Int32:
		typeName = "int32"
	case reflect.String:
		typeName = "string"
	// ... handle more cases
	default:
		typeName = "unrecognized type"
	}
	fmt.Printf("v is of type '%s'. Size: %d bytes\n", typeName, typ.Size())
}

// :show end

func main() {
	// :show start
	printType(int32(3))
	printType("")
	i := 3
	printType(&i) // *int i.e. pointer to int
	// :show end
}
