package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
)

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

// :show start
func decodeFromReader(r io.Reader) (*People, error) {
	var people People
	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&people)
	if err != nil {
		return nil, err
	}
	return &people, nil
}

func decodeFromString(s string) (*People, error) {
	r := bytes.NewBufferString(s)
	return decodeFromReader(r)
}

func decodeFromFile(path string) (*People, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return decodeFromReader(f)
}

// :show end

func main() {
	people, err := decodeFromString(xmlStr)
	if err != nil {
		log.Fatalf("decodeFromString failed with '%s'\n", err)
	}
	fmt.Printf("%#v\n\n", people)
}
