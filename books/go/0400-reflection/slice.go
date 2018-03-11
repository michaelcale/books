package main

import (
	"fmt"
	"reflect"
)

func main() {
	// :show start
	a := []int{3, 1, 8}
	rv := reflect.ValueOf(a)

	fmt.Printf("len(a): %d\n", rv.Len())
	fmt.Printf("cap(a): %d\n", rv.Cap())

	fmt.Printf("slice kind: '%s'\n", rv.Kind().String())

	fmt.Printf("element type: '%s'\n", rv.Type().Elem().Name())

	el := rv.Index(0).Interface()
	fmt.Printf("a[0]: %v\n", el)

	elRef := rv.Index(1)
	fmt.Printf("elRef.CanAddr(): %v\n", elRef.CanAddr())
	fmt.Printf("elRef.CanSet(): %v\n", elRef.CanSet())

	elRef.SetInt(5)
	fmt.Printf("a: %v\n", a)

	// :show end
}
