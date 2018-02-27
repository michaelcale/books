package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
type Person struct {
	fullName string
	Name     string
	Age      int    `xml:"age"`
	City     string `xml:"city"`
}

// :show end

func main() {
	// :show start
	p := Person{
		Name: "John",
		Age:  37,
		City: "SF",
	}
	d, err := xml.Marshal(&p)
	if err != nil {
		log.Fatalf("xml.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Person in compact XML: %s\n", string(d))

	d, err = xml.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf("xml.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Person in pretty-printed XML:\n%s\n", string(d))
	// :show end
}
