---
Title: JSON
Id: 994
---
## Syntax
- func Marshal(v interface{}) ([]byte, error)
- func Unmarshal(data []byte, v interface{}) error

## Remarks
The package [`"encoding/json"`](https://golang.org/pkg/encoding/json/) Package json implements encoding and decoding of JSON objects in `Go`.

----------

Types in JSON along with their corresponding concrete types in Go are:

| JSON Type | Go Concrete Type |
| ------ | ------ |
| boolean   | bool   |
| numbers   | float64 or int   |
| string   | string   |
| null   | nil   |
