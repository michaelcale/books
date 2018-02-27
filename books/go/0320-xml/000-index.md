---
Title: XML
Id: 1846
---
Package [`encoding/xml`](https://godoc.org/encoding/xml) in Go standard library provides functionality for serializing data as XML and parsing XML.

## Serialize a struct as XML

@file index.go output sha1:b79884448d22013646eb2c6a46097d188c9ac10c goplayground:Klo84laKIF0

Both `xml.Marshal` and `xml.MarshalIndent` take [`interface{}`](a-der300hf) as first argument. We can pass any Go value and it'll be wrapped into `interface{}` with their type.

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

## Parse XML into a Go struct

@file index2.go output sha1:e08744bc021e4d568e1e7ffe83533b4b292039cf goplayground:6TeIISoehvE

Parsing is the opposite of serializing.

Unlike with serializing, when parsing into structs, we must pass a pointer to struct. Otherwise `xml.Unmarshal` will receive and modify a copy of the struct, not the struct itself. The copy will be then discarded after returning from `xml.Unmarshal`.

Notice that XML element `city` was decoded into `City` struct field even though the names don't match and we didn't provide explicit mapping with `xml` struct tag.

That happened because XML decoder has a little bit of smarts when matching dictionary key names to struct field names. It's best to not rely on such smarts and define mappings explicitly.

All struct fields are optional and when not present in XML text their values will be untouched. When decoding into newly initialized struct their value will be zero value for a given type.

Field `Name` shows that XML decoder can also automatically decode into a pointer to a value.

This is useful when you need to know if a value was present in XML or not. If we used `string` for `Name`, we wouldn't know if value of empty string means that XML had `name` key with empty string as a value or is it because the value wasn't there at all.

By using a pointer to a string we know that `nil` means there was no value.
