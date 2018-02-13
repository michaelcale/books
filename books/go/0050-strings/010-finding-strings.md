---
Title: Find string in another string
Search: searching for strings
Id: 30331
---

## Find position of string in another string

Using [`strings.Index`](https://golang.org/pkg/strings/#Index):

@file finding_strings.go output

## Find position of string in another string from end

Above we searched from the beginning of the string. We can also search from end using [`strings.LastIndex`](https://golang.org/pkg/strings/#LastIndex):

@file finding_strings2.go output
```

## Find all occurences of a substring

Above we only found first occurence of substring. Here's how to find all of them:

@file finding_strings3.go output

<!-- TODO: example using a regular expression -->

## Check if a string contains another string

Using [`strings.Contains`](https://golang.org/pkg/strings/#Contains):

@file finding_strings4.go output

## Check if a strings starts with another strings

Using [`strings.HasPrefix`](https://golang.org/pkg/strings/#HasPrefix):

@file finding_strings5.go output

## Check if a strings ends with another strings

Using [`strings.HasSuffix`](https://golang.org/pkg/strings/#HasSuffix):

@file finding_strings5.go output
