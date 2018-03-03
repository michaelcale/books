package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// :show start
type Student struct {
	Name     string
	Standard int `json:"Standard"`
}

// :show end

func main() {
	// :show start
	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var dec *json.Decoder
	dec = json.NewDecoder(f)
	if err != nil {
		log.Fatal(err)
	}

	var studentList []Student

	err = dec.Decode(&studentList)
	if err != nil {
		log.Fatal(err)
	}

	for _, student := range studentList {
		fmt.Printf("Student: %s, standard: %d\n", student.Name, student.Standard)
	}
	// :show end
}
