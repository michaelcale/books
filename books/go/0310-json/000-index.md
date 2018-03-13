---
Title: JSON
Id: 182
SOId: 994
---

Package [`encoding/json`](https://golang.org/pkg/encoding/json/) in standard library provides functionality for serializing data as JSON and parsing JSON into data.

## Serialize a struct as JSON

@file json_serialize.go output sha1:a91894315e1a06003e4862d5fbc12003cfcd20fd goplayground:T6_VXCMeO3r

Both `json.Marshal` and `json.MarshalIndent` take [`interface{}`](94) as first argument. We can pass any Go value, it'll be wrapped into `interface{}` with their type.

Marshaller will use reflection to inspect passed value and encode it as JSON strings.

When serializing structs, only exported fields (whose names start with capital letter) are serialized / deserialized.

In our example, `fullName` is not serialized.

Structs are serialized as JSON dictionaries. By default dictionary keys are the same as struct field names.

Struct field `Name` is serialized under dictionary key `Name`.

We can provide custom mappings with struct tags.

We can attach arbitrary struct tags string to struct fields.

`json:"age"` instructs JSON encoder / decoder to use name `age` for dictionary key representing field `Age`.

When serializing structs, passing the value and a pointer to it generates the same result.

Passing a pointer is more efficient because passing by value creates unnecessary copy.

`json.MarshallIndent` allows for pretty-printing of nested structures. The result takes up more space but is easier to read.

## Parse JSON into a struct

@file json_deserialize.go output sha1:1adbbe8fa6abf4c08bf821b2b3e7db1c82ace85d goplayground:CllMvV6twjM

Parsing is the opposite of serializing.

Unlike with serializing, when parsing into structs, we must pass a pointer to struct. Otherwise `json.Unmarshal` will receive and modify a copy of the struct, not the struct itself. The copy will be then discarded after returning from `json.Unmarshal`.

Notice that JSON element `city` was decoded into `City` struct field even though the names don't match and we didn't provide explicit mapping with `json` struct tag.

That happened because JSON decoder has a little bit of smarts when matching dictionary key names to struct field names. It's best to not rely on such smarts and define mappings explicitly.

All struct fields are optional and when not present in JSON text their values will be untouched. When decoding into newly initialized struct their value will be zero value for a given type.

Field `Name` shows that JSON decoder can also automatically decode into a pointer to a value.

This is useful when you need to know if a value was present in JSON or not. If we used `string` for `Name`, we wouldn't know if value of empty string means that JSON had `name` key with empty string as a value or is it because the value wasn't there at all.

By using a pointer to a string we know that `nil` means there was no value.

## Go to JSON type mapping

| JSON Type  | Go Concrete Type |
| ---------- | ---------------- |
| boolean    | bool             |
| numbers    | float64 or int   |
| string     | string           |
| array      | slice            |
| dictionary | struct           |
| null       | nil              |

See more in [type mappings](188).
