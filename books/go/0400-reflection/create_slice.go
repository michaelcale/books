package main

import (
	"fmt"
	"reflect"
)

func main() {
	// :show start
	typ := reflect.SliceOf(reflect.TypeOf("example"))
	// create slice with capacity 10 and length 1
	rv := reflect.MakeSlice(typ, 1, 10)
	rv.Index(0).SetString("foo")

	a := rv.Interface().([]string)
	fmt.Printf("a: %#v\n", a)
	// :show end
}
