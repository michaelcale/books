package main

import (
	"fmt"
	"reflect"
)

// :show start

type S struct {
	FirstName string `my_tag:"first-name"`
	lastName  string
	Age       int `json:"age",xml:"AgeXml`
}

func describeStructSimple(rv reflect.Value) {
	structType := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		v := rv.Field(i)
		structField := structType.Field(i)
		name := structField.Name
		typ := structField.Type
		tag := structField.Tag
		jsonTag := tag.Get("json")
		isExported := structField.PkgPath == ""
		if isExported {
			fmt.Printf("name: '%s',\ttype: '%s', value: %v,\ttag: '%s',\tjson tag: '%s'\n", name, typ, v.Interface(), tag, jsonTag)
		} else {
			fmt.Printf("name: '%s',\ttype: '%s',\tvalue: not accessible\n", name, v.Type().Name())
		}
	}
}

func main() {
	s := S{
		FirstName: "John",
		lastName:  "Doe",
		Age:       27,
	}
	describeStructSimple(reflect.ValueOf(s))
}

// :show end
