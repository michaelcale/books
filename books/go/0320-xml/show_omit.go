package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
type ShowOmit struct {
	Name          string `xml:"name"`
	NotSerialized string `xml:"-"`
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
	v := &ShowOmit{
		Name:          "John",
		NotSerialized: "Connor",
	}
	printXML(v)
	// :show end
}
