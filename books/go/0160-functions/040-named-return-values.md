Title: Named return values
Id: 1253
Body:
Return values can be assigned to a local variable.

An empty `return` statement can then be used to return their current values. This is known as *"naked"* return.

Naked return statements should be used only in short functions as they harm readability in longer functions:

@file named_return_values.go output

Two important things must be noted:

- parenthesis around the return values are **mandatory**.
- empty `return` statement must always be provided.
