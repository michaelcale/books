---
Title: Zero values
Id: 6069
---
Variables in Go are always initialized with a known value even if not explicitly assigned in source code.

Each Go type has a zero value.

Variables that are not explicitly initialized (assigned an explicit value) will have a value equal to the zero value for their type.

This is different from C/C++, where variables that are not explicitly assigned have undefined values.

The values of zero type are unsurprising:

|type|zero value|
|----|----------|
|bool|false|
|integers|0|
|floating poing numbers|0.0|
|string|""|
|pointer|nil|
|slice|nil|
|map|nil|
|interface|nil|
|channel|nil|
|array|all elements have zero values|
|struct|all members set to zero value of their type|
|function|nil|

Said differently:

@file zero_values.go output sha1:7bed196df8157c429f9c3ad7801b0e69f53485e5 goplayground:9B0IqtjU0RF

