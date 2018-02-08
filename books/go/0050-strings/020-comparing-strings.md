---
Title: Comparing strings
Id: 29828
---

## Compare strings with ==, > and <

``go
s1 := "string one"
s2 := "string two"

if s1 == s2 {
    fmt.Printf("s1 is equal to s2\n")
} else {
    fmt.Printf("s1 is not equal to s2\n")
}

if s1 == s1 {
    fmt.Printf("s1 is equal to s1\n")
} else {
    fmt.Printf("inconcivable! s1 is not equal to itself\n")
}

if s1 > s2 {
    fmt.Printf("s1 is > than s2\n")
} else {
    fmt.Printf("s1 is not > than s2\n")
}

if s1 < s2 {
    fmt.Printf("s1 is < than s2\n")
} else {
    fmt.Printf("s1 is not <> than s2\n")
}
```

Comparison is performed on raw bytes.

This works as you would expect for ascii (i.e. english) text but might not be what you mean when strings used mixed case (e.g. "abba" is > "Zorro") or use letters from non-english alphabets.

## Compare with `strings.Compare`

You can also compare with [strings.Compare](https://golang.org/pkg/strings/#Compare) but use `==`, `>` and `>` instead as it has the same semantics.

## Case-insensitive compare

[strings.EqualFold](https://golang.org/pkg/strings/#EqualFold)

Sometimes you want "Go" to equal "go", which is not the case when using `==`.

```go
s1 := "gone"
s2 := "GoNe"
if strings.EqualFold(s1, s2) {
    fmt.Printf("'%s' is equal '%s' when ignoring case\n", s1, s2)
} else {
    fmt.Printf("'%s' is not equal '%s' when ignoring case\n", s1, s2)
}
```

The exact rule is: both string are considered UTF-8-encoded strings and characters are compared using Unicode case-folding.

