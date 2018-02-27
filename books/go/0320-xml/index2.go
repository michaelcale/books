package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
type Person struct {
	Name       *string `xml:"name"`
	Age        int     `xml:"age"`
	City       string
	Occupation string
}

var xmlStr = `<Person>
	<name>John</name>
	<age>37</age>
	<city>SF</city>
</Person>
`

// :show end

func main() {
	// :show start
	var p Person
	err := xml.Unmarshal([]byte(xmlStr), &p)
	if err != nil {
		log.Fatalf("xml.Unmarshal failed with '%s'\n", err)
	}
	fmt.Printf("Person struct parsed from XML: %#v\n", p)
	fmt.Printf("Name: %#v\n", *p.Name)
	// :show end
}
