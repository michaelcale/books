---
Title: String
Id: 9666
---
A string is in effect a read-only slice of bytes. In Go a string literal will always contain a valid UTF-8 representation of its content.

## Syntax
- variableName := "Hello World" // declare a string
- variableName := \`Hello World` // declare a raw literal string
- variableName := "Hello " + "World" // concatenates strings
- substring := "Hello World"[0:4] // get a part of the string
- letter := "Hello World"[6] // get a character of the string
- fmt.Sprintf("%s", "Hello World") // formats a string
