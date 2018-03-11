package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start

type ShowOmitEmpty struct {
	NonEmpty string `xml:",omitempty"`
	Empty    string `xml:",omitempty"`
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
	v := &ShowOmitEmpty{
		NonEmpty: "non empty",
		Empty:    "",
	}
	printXML(v)

	// :show end
}
