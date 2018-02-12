---
Title: Blank identifier
Id: 29103
---
To help avoid mistakes, Go compiler doesn't allow unused variables.

However, there are some situations when you don't need to use a value stored in a variable.

In those cases, you use a "blank identifier" `_` to assign and discard the assigned value.

A blank identifier can be assigned a value of any type, and is most commonly used in functions that return multiple values.

**Multiple Return Values**

@file blank_identifier.go output

**Using `range`**

@file blank_identifier_2.go output

