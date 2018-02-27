---
Title: XML
Id: 1846
---
Package [`encoding/xml`](https://godoc.org/encoding/xml) in Go standard library provides functionality for serializing data as XML and parsing XML.

## Serialize a struct as XML

@file index.go output

Both `xml.Marshal` and `xml.MarshalIndent` take [`interface{}`](ch-der300hf) as first argument. We can pass any Go value and it'll be wrapped into `interface{}` with their type.

Marshaller will use reflection to to inspect passed value and encode it as XML strings.

When serializing structs, only exported fields (whose names start with capital letter) are serialized / deserialized.

In our example, `fullName` is not serialized.

Structs are serialized as XML elements. By default element names are the same as struct field names.

Struct field `Name` is serialized under dictionary key `Name`.

We can provide custom mappings with struct tags.

We can attach arbitrary struct tags string to struct fields.

`xml:"age"` instructs XML encoder / decoder to use name `age` for element name representing field `Age`.

When serializing structs, passing the value and a pointer to it generates the same result.

Passing a pointer is more efficient becase passing by value creates unnecessary copy.

`xml.MarshallIndent` allows for pretty-printing of nested structures. It's less efficient but the result is easier for humans to read.
