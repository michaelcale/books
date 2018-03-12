package main

import (
	"fmt"
	"reflect"
)

// :show start
type S struct {
	N int
}

func setIntPtr() {
	var n int = 2
	reflect.ValueOf(&n).Elem().SetInt(4)
	fmt.Printf("setIntPtr: n=%d\n", n)
}

func setStructFieldDirect() {
	var s S
	reflect.ValueOf(&s.N).Elem().SetInt(5)
	fmt.Printf("setStructFieldDirect: n=%d\n", s.N)
}

func setStructPtrField() {
	var s S
	reflect.ValueOf(&s).Elem().Field(0).SetInt(6)
	fmt.Printf("setStructPtrField: s.N: %d\n", s.N)
}

func handlePanic(funcName string) {
	if msg := recover(); msg != nil {
		fmt.Printf("%s panicked with '%s'\n", funcName, msg)
	}
}

func setStructField() {
	defer handlePanic("setStructField")
	var s S
	reflect.ValueOf(s).Elem().Field(0).SetInt(4)
	fmt.Printf("s.N: %d\n", s.N)
}

func setInt() {
	defer handlePanic("setInt")
	var n int = 2
	rv := reflect.ValueOf(n)
	rv.Elem().SetInt(4)
}

func setIntPtrWithString() {
	defer handlePanic("setIntPtrWithString")
	var n int = 2
	reflect.ValueOf(&n).Elem().SetString("8")
}

// :show end

func main() {
	setIntPtr()
	setStructFieldDirect()
	setStructPtrField()

	setInt()
	setStructField()

	setIntPtrWithString()
}
