package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
type ShowAttr struct {
	B bool `xml:"n,attr"`
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
	v := &ShowAttr{
		B: true,
	}
	printXML(v)

	// :show end
}
