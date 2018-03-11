---
Title: Named return values
Id: 115
SOId: 1253
---

Return values can be assigned to a local variable.

An empty `return` statement can then be used to return their current values. This is known as _"naked"_ return.

Naked return statements should be used only in short functions as they harm readability in longer functions:

@file named_return_values.go output sha1:02946dc012fe64e4490cdc382b63d9024efe929a goplayground:s3zX6fDUijB

Two important things must be noted:

* parenthesis around the return values are **mandatory**.
* empty `return` statement must always be provided.
