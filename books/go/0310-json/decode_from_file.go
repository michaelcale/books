package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// :show start
type Student struct {
	Name     string
	Standard int `json:"Standard"`
}

func decodeFromReader(r io.Reader) ([]*Student, error) {
	var res []*Student

	dec := json.NewDecoder(r)
	err := dec.Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func decodeFromString(s string) ([]*Student, error) {
	r := bytes.NewBufferString(s)
	return decodeFromReader(r)
}

func decodeFromFile(path string) ([]*Student, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decodeFromReader(f)
}

// :show end

const jsonStr = `
[
    {
        "Name" : "John Doe",
        "Standard" : 4
    },
    {
        "Name" : "Peter Parker",
        "Standard" : 11
    },
    {
        "Name" : "Bilbo Baggins",
        "Standard" : 150
    }
]
`

func main() {
	// studentList, err := decodeFromFile("data.json")
	studentList, err := decodeFromString(jsonStr)
	if err != nil {
		log.Fatalf("decodeFromString() failed with '%s'\n", err)
	}
	for _, student := range studentList {
		fmt.Printf("Student: %s, standard: %d\n", student.Name, student.Standard)
	}
}
