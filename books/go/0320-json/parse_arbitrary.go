package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// :show start
var jsonStr = `{
	"name": "Jane",
	"age": 24,
	"city": "ny"
}`

// :show end

func main() {
	// :show start
	var doc map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &doc)
	if err != nil {
		log.Fatalf("json.Unmarshal failed with '%s'\n", err)
	}
	fmt.Printf("doc: %#v\n", doc)
	name, ok := doc["name"].(string)
	if !ok {
		log.Fatalf("doc has no key 'name' or its value is not string\n")
	}
	fmt.Printf("name: %#v\n", name)
	// :show end
}
