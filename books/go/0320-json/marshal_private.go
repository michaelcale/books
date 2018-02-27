package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// :show start
type MyStruct struct {
	uuid string
	Name string
}

func (m MyStruct) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Uuid string
		Name string
	}{
		Uuid: m.uuid,
		Name: m.Name,
	})
	if err != nil {
		return nil, err
	}
	return j, nil
}

// :show end

func main() {
	// :show start
	s := MyStruct{
		uuid: "uid-john",
		Name: "John",
	}
	d, err := json.Marshal(&s)
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Person in compact JSON: %s\n", string(d))

	// :show end
}
