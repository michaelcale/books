package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start

type ShowXMLName struct {
	XMLName xml.Name `xml:"data"`
	N       int      `xml:"n"`
}

// :show end

func printXML(v interface{}) {
	d, err := xml.Marshal(v)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("XML: %s\n\n", string(d))
}

func main() {
	// :show start
	v := &ShowXMLName{
		N: 5,
	}
	printXML(v)

	// :show end
}
