// :run go run -race $file
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
	fmt.Printf("Type '%s', its size: %d bytes\n", typ.Name(), typ.Size())
	// :show end
}
