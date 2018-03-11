---
Title: XML
Id: 189
---

Package [`encoding/xml`](https://godoc.org/encoding/xml) in standard library provides functionality for serializing data as XML and parsing XML.

## Parse XML into a struct

Parsing XML is similar to parsing JSON. You define structures that map to the structure of XML and unmarshal from `[]byte` slice or `io.Reader` into a struct.

@file unmarshal_simple.go output sha1:32254e64f0145fbc199c82613444ac313a54beae goplayground:1nB6hJ_EuU2

Unlike with serializing, when unmarshalling into a struct, we must pass a pointer to struct. Otherwise `xml.Unmarshal` will receive and modify a copy of the struct, not the struct itself. The copy will be then discarded after returning from `xml.Unmarshal`.

In XML a value can be represented as element (`<state>CA</state>`) or attribute (`<person age="34">`).

In xml struct tags element is the default. To switch to attribute, add `,attr` to struct xml tag as done for `Age`.

All struct fields are optional. When not present in XML text their values will be untouched. When decoding into newly initialized struct their value will be zero value for a given type.

Field `City` shows that XML decoder can automatically decode into a pointer to a value.

This is useful when you need to know if a value was present in XML or not. If we used `string` for `City` field, we wouldn't know if empty string means:

* `city` element was not present in XML
* element was present but had empty value (`<city></city>`)

By using a pointer to a string we know that `nil` means there was no value.

## Serialize a struct as XML

@file marshal_simple.go output sha1:399c70c8b4462adfef0eb8cdf0a52fe7c2c1e43c goplayground:U3BoXM0mqnI

Both `xml.Marshal` and `xml.MarshalIndent` take [`interface{}`](94) as first argument. We can pass any Go value and it'll be wrapped into `interface{}` with their type.

Marshaller will use [reflection](1854) to inspect passed value and encode it as XML strings.

When serializing structs, only exported fields (whose names start with capital letter) are serialized / deserialized.

In our example, `noteSerialized` is not serialized.

Structs are serialized as XML elements.

Simple types can be serialized as XML elements or attributes (`Age` field in `Person` struct)

By default name of element / attribute is the same as name of the field.

We can provide custom mappings with struct tags.

`xml:"city"` tells XML encoder to use name `city` for field `City`.

When serializing structs, passing the value or a pointer to `xml.Marshal` generates the same result.

Passing a pointer is more efficient becase passing by value creates unnecessary copy.

`xml.MarshallIndent` allows for pretty-printing of nested structures. The result takes up more space but is easier to read.

`XMLName` allows to control the name of top-level element. In our example, without providing XML element, the data would be serialized based on the struct name:

```xml
<People>
    ....
</People>
```
