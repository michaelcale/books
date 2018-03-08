---
Title: Parsing arbitrary JSON documents
Id: 4113
---
Parsing [into a struct](994) is very convenient but sometimes you don't know the structure of JSON document upfront.

For arbitrary JSON documents we can decode into a `map[string]interface{}`, which can represent all valid JSON documents.

@file parse_arbitrary.go output sha1:7ac4d0f515cd1b43a884374e86a6689b5a943ee0 goplayground:rN9W0rFQ8sN

For basic JSON types, the value in the map is `bool`, `int`, `float64` or `string`.

For JSON arrays, the value is `[]interface{}`.

For JSON dictionaries, the value is (again) `map[string]interface{}`.

This approach is flexible but dealing with `map[string]interface{}` to access values is rather painful.
