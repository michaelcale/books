package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
type ShowInnerXML struct {
	Str string `xml:"s"`
	Raw string `xml:",innerxml"`
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
	v := &ShowInnerXML{
		Str: "<foo></foo>",
		Raw: "<foo></foo>",
	}
	printXML(v)

	// :show end
}
