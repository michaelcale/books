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
	<!-- sample comment -->
	<person age="23">
		<first-name>Julia</first-name>
	</person>
</people>`

// :show end

func main() {
	// :show start
	r := bytes.NewBufferString(xmlStr)
	decoder := xml.NewDecoder(r)
	inCityElement := false
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
			if v.Name.Local == "person" {
				for _, attr := range v.Attr {
					if attr.Name.Local == "age" {
						fmt.Printf("Element: '<person>', attribute 'age' has value '%s'\n", attr.Value)
					}
				}
			} else if v.Name.Local == "city" {
				inCityElement = true
			}

		case xml.EndElement:
			if v.Name.Local == "city" {
				inCityElement = false
			}

		case xml.CharData:
			if inCityElement {
				fmt.Printf("City: %s\n", string(v))
			}

		case xml.Comment:
			fmt.Printf("Comment: %s\n", string(v))

		case xml.ProcInst:
			// handle XML processing instruction like <?target inst?>

		case xml.Directive:
			// handle XML directive like <!text>
		}
	}
	// :show end
}
