package main

import (
	"fmt"
)

// :show start
// User describes a user
type User struct {
	FirstName, LastName string
	Email               string
	Age                 int
	userID              int
}

// FullName returns full name of a user
func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func main() {
	// zero value of struct
	var u User
	fmt.Printf("u: %#v\n\n", u)

	// pu is *User i.e. a pointer to User struct
	pu := new(User)
	pu.Age = 33
	fmt.Printf("*pu: %#v\n", *pu)

	// &User{} is the same as new(User)
	pu = &User{}
	pu.Age = 18
	fmt.Printf("*pu: %#v\n", *pu)

	pu.FirstName, pu.LastName = "Jane", "Doe"
	fmt.Printf("pu.FullName(): %s\n", pu.FullName())
}

// :show end
