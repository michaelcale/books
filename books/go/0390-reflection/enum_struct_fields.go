package main

import (
	"fmt"
	"reflect"
	"strings"
)

// :show start
type Inner struct {
	N int
}

type S struct {
	Inner
	NamedInner Inner
	PtrInner   *Inner
	unexported int
	N          int8
}

func indentStr(level int) string {
	return strings.Repeat("  ", level)
}

// if sf is not nil, this is a field of a struct
func describeStruct(level int, rv reflect.Value, sf *reflect.StructField) {
	structType := rv.Type()
	nFields := rv.NumField()
	typ := rv.Type()
	if sf == nil {
		fmt.Printf("%sstruct %s, %d field(s), size: %d bytes\n", indentStr(level), structType.Name(), nFields, typ.Size())
	} else {
		fmt.Printf("%sname: '%s' type: 'struct %s', offset: %d, %d field(s), size: %d bytes, embedded: %v\n", indentStr(level), sf.Name, structType.Name(), sf.Offset, nFields, typ.Size(), sf.Anonymous)
	}

	for i := 0; i < nFields; i++ {
		fv := rv.Field(i)
		sf := structType.Field(i)
		describeType(level+1, fv, &sf)
	}
}

// if sf is not nil, this is a field of a struct
func describeType(level int, rv reflect.Value, sf *reflect.StructField) {
	switch rv.Kind() {

	case reflect.Int, reflect.Int8:
		// in real code we would handle more primitive types
		i := rv.Int()
		typ := rv.Type()
		if sf == nil {
			fmt.Printf("%stype: '%s', value: '%d'\n", indentStr(level), typ.Name(), i)
		} else {
			fmt.Printf("%s name: '%s' type: '%s', value: '%d', offset: %d, size: %d\n", indentStr(level), sf.Name, typ.Name(), i, sf.Offset, typ.Size())
		}

	case reflect.Ptr:
		fmt.Printf("%spointer\n", indentStr(level))
		describeType(level+1, rv.Elem(), nil)

	case reflect.Struct:
		describeStruct(level, rv, sf)
	}
}

func main() {
	var s S
	describeType(0, reflect.ValueOf(s), nil)
}

// :show end
