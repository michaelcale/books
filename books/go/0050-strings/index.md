---
Title: Strings
Id: 9666
---

Strings in Go are immutable sequences of bytes.

Unlike languages like Python or Java, they are not internally represented as Unicode. Consequently, when reading strings from files or network connections, there is no conversion step from bytes to internal representation. When writing strings to files, there is no conversion to a code page.

Go strings don't assume any particular code page. They are just bytes.

Go source code files are always UTF-8 so strings defined in source code are also valid utf8 strings.

Additionally, functions in standard library that involve converting characters to upper-case or lower-case etc. assume that raw bytes represent UTF-8-encoded Unicode strings and perform transformations using Unicode rules.

Basic string usage:
```go
var s string // empty string ""
s1 := "string\nliteral\nwith\tescape characters"
s2 := `raw string literal
which doesn't recgonize escape characters like \n
`

// you can add strings with +
fmt.Printf("sum of string: %s\n", s + s1 + s2)

// you can compare strings with ==
if s1 == s2 {
    fmt.Pritnf("s1 is equal to s2\n")
} else {
    fmt.Pritnf("s1 is not equal to s2\n")
}

fmt.Printf("substring of s1: %s\n", s1[3:5])
fmt.Printf("byte (character) at position 3 in s1: %d\n", s1[3])

// C-style string formatting
s = fmt.Sprintf("%d + %f = %s", 1, float64(3), "4")
fmt.Printf("s: %s\n", s)
```

Important standard library packages for working on strings:
* [strings](https://golang.org/pkg/strings/) implements string searching, splitting, case conversions
* [bytes](https://golang.org/pkg/bytes/) has the same functionality as `strings` package but operating on `[]byte` byte slices
* [strconv](https://golang.org/pkg/strconv/) for conversion between strings and integer and float numbers
* [unicode/utf8](https://golang.org/pkg/unicode/utf8/) decodes from UTF-8-encoded strings and encodes into UTF-8-encoded string
* [regexp](https://golang.org/pkg/regexp/) implements regular expressions
* [text/scanner](https://golang.org/pkg/text/scanner/) for scanning and tokenizing UTF-8-encoded text
* [text/template](https://golang.org/pkg/text/template/) for generating larger strings from templates
* [html/template](https://golang.org/pkg/html/template/) has all the functionaliy of `text/template` but understoods the structure of HTML for generation of HTML safe from code injection attacks
