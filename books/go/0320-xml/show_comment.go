package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

// :show start
type ShowComment struct {
	Str string `xml:",comment"`
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
	v := &ShowComment{
		Str: "comment",
	}
	printXML(v)

	// :show end
}
