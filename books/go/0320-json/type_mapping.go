package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// :show start
func printSerialized(v interface{}) {
	d, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("json.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("%T: %s\n", v, string(d))
}

// :show end

func main() {
	// :show start
	printSerialized(nil)
	printSerialized(5)
	printSerialized(8.23)
	printSerialized("john")
	ai := []int{5, 4, 18}
	printSerialized(ai)
	a := []interface{}{4, "string"}
	printSerialized(a)
	d := map[string]interface{}{
		"i": 5,
		"s": "foo",
	}
	printSerialized(d)
	s := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  37,
	}
	printSerialized(s)
	// :show end
}
