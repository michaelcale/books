package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
var xmlStr = `
<people>
	<person age="34">
		<first-name>John</first-name>
		<address>
			<city>San Francisco</city>
			<state>CA</state>
		</address>
	</person>

	<person age="23">
		<first-name>Julia</first-name>
	</person>
</people>`

type People struct {
	Person []Person `xml:"person"`
}

type Person struct {
	Age       int     `xml:"age,attr"`
	FirstName string  `xml:"first-name"`
	Address   Address `xml:"address"`
}

type Address struct {
	City  *string `xml:"city"`
	State string  `xml:"state"`
}

// :show end

func main() {
	// :show start
	var people People
	data := []byte(xmlStr)
	err := xml.Unmarshal(data, &people)
	if err != nil {
		log.Fatalf("xml.Unmarshal failed with '%s'\n", err)
	}
	fmt.Printf("%#v\n\n", people)
	// :show end
}
