---
Title: Basic declaration
Id: 1254
---
A basic struct is declared as follows:

@file declaring_structs.go output

Each value is called a field.

Fields are usually written one per line, with the field's name preceeding its type.

Consecutive fields of the same type may be combined, as `FirstName` and `LastName` in the above example.

Field names that start with upper case (`FirstName`, `Email`) are public i.e. accesible by all code.

Field names that start with lower case (`userID`) are private and only accessible by code in the same package.

