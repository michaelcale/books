package main

import (
	"fmt"
	"strconv"
)

// :show start
// Stringer is an interface with a single method
type Stringer interface {
	String() string
}

// User struct that implements Stringer interface
type User struct {
	Name string
}

func (u *User) String() string {
	return u.Name
}

// Any type can implement an interface. Here we create
// an alias of int type an implement Stringer interface

type MyInt int

func (mi MyInt) String() string {
	return strconv.Itoa(int(mi))
}

// printTypeAndString accepts an interface. 's' can be any value
// that implements Stringer interface
func printTypeAndString(s Stringer) {
	fmt.Printf("%T: %s\n", s, s)
}

func main() {
	u := &User{Name: "John"}
	printTypeAndString(u)

	n := MyInt(5)
	printTypeAndString(n)
}

// :show end
