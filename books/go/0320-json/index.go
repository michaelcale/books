package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// :show start
type Person struct {
	fullName string
	Name     string
	Age      int    `json:"age"`
	City     string `json:"city"`
}

// :show end

func main() {
	// :show start
	p := Person{
		Name: "John",
		Age:  37,
		City: "SF",
	}
	d, err := json.Marshal(&p)
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Person in compact JSON: %s\n", string(d))

	d, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Person in pretty-printed JSON:\n%s\n", string(d))
	// :show end
}
