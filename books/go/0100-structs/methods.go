package main

import "fmt"

// :show start

type User struct {
	name string
}

func (u User) Name() string {
	return u.name
}

func (u *User) SetName(newName string) {
	u.name = newName
}

func main() {
	var me User

	me.SetName("Slim Shady")
	fmt.Println("My name is", me.Name())
}

// :show end
