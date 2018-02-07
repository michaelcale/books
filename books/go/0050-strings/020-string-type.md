---
Title: String type
Id: 29828
---
The `string` type allows you to store text, which is a series of characters. There are multiple ways to create strings. A literal string is created by writing the text between double quotes.

    text := "Hello World"

Because Go strings support UTF-8, the previous example is perfectly valid. Strings hold arbitrary bytes which does not necessarily mean every string will contain valid UTF-8 but string literals will always hold valid UTF-8 sequences.

The zero value of strings is an empty string `""`.

Strings can be concatenated using the `+` operator.

    text := "Hello " + "World"

Strings can also be defined using backticks ` `` `. This creates a raw string literal which means characters won't be escaped.

```go
text1 := "Hello\nWorld"
text2 := `Hello
World`
```

In the previous example, `text1` escapes the `\n` character which represents a new line while `text2` contains the new line character directly. If you compare `text1 == text2` the result will be `true`.

However, ``text2 := `Hello\nWorld` `` would not escape the `\n` character which means the string contains the text `Hello\nWorld` without a new line. It would be the equivalent of typing `text1 := "Hello\\nWorld"`.
