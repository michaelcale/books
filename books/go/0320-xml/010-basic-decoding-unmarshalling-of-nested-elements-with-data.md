---
Title: Basic decoding / unmarshalling of nested elements with data
Id: 190
Score: 1
SOId: 6046
---
XML elements often nest, have data in attributes and/or as character data. The way to capture this data is by using `,attr` and `,chardata` respectively for those cases.

```go
var doc = `
<parent>
  <child1 attr1="attribute one"/>
  <child2>and some cdata</child2>
</parent>
`

type parent struct {
    Child1 child1 `xml:"child1"`
    Child2 child2 `xml:"child2"`
}

type child1 struct {
    Attr1 string `xml:"attr1,attr"`
}

type child2 struct {
    Cdata1 string `xml:",cdata"`
}

func main() {
    var obj parent
    err := xml.Unmarshal([]byte(doc), &obj)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(obj.Child2.Cdata1)
}
```
