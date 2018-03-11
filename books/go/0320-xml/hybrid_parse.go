package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
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
		<address>
			<city>Austin</city>
			<state>TX</state>
		</address>
	</person>
</people>`

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

// :show end

func main() {
	// :show start
	r := bytes.NewBufferString(xmlStr)
	decoder := xml.NewDecoder(r)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			// io.EOF is a successful end
			break
		}
		if err != nil {
			fmt.Printf("decoder.Token() failed with '%s'\n", err)
			break
		}

		switch v := t.(type) {

		case xml.StartElement:
			if v.Name.Local == "address" {
				var address Address
				err = decoder.DecodeElement(&address, &v)
				if err != nil {
					fmt.Printf("decoder.DecodeElement() failed with '%s'\n", err)
					break
				}
				fmt.Printf("%+#v\n", address)
			}
		}
	}
	// :show end
}
