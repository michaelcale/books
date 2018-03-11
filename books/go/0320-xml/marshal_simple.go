package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
type People struct {
	XMLName        xml.Name `xml:"people"`
	Person         []Person `xml:"person"`
	noteSerialized int
}

type Person struct {
	Age       int     `xml:"age,attr"`
	FirstName string  `xml:"first-name"`
	Address   Address `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

// :show end

func main() {
	// :show start
	people := People{
		Person: []Person{
			Person{
				Age:       34,
				FirstName: "John",
				Address:   Address{City: "San Francisco", State: "CA"},
			},
		},
		noteSerialized: 8,
	}
	d, err := xml.Marshal(&people)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("Compact XML: %s\n\n", string(d))

	d, err = xml.MarshalIndent(&people, "", "  ")
	if err != nil {
		log.Fatalf("xml.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Pretty printed XML:\n%s\n", string(d))

	// :show end
}
