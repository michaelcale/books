package main

import (
	"context"
	"fmt"
)

// :show start
// User describes a user
type User struct {
	Name       string
	IsLoggedIn bool
}

type userKeyType int

var userKey userKeyType

func contextWithUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

// returns nil if not set
func getUserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(userKey).(*User)
	if !ok {
		return nil
	}
	return user
}

// will panic if not set
func mustGetUserFromContext(ctx context.Context) *User {
	return ctx.Value(userKey).(*User)
}

func printUser(ctx context.Context) {
	user := getUserFromContext(ctx)
	fmt.Printf("User: %#v\n", user)
}

func main() {
	ctx := context.Background()
	user := &User{
		Name:       "John",
		IsLoggedIn: false,
	}
	ctx = contextWithUser(ctx, user)

	printUser(ctx)
}

// :show end
