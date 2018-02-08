---
Title: Finding strings inside strings
Search: searching for strings
Id: 30331
Score: 1
---

## Find position of string in another string

Using [`strings.Index`](https://golang.org/pkg/strings/#Index):

```go
s := "where hello is?"
toFind := "hello"
idx := strings.Index(s, toFind)
fmt.Printf("'%s' is in s starting at position %d\n", toFind, idx)

// when string is not found, result is -1
idx = strings.Index(s, "not present")
fmt.Printf("Index of non-existent substring is: %d\n", idx)
```

## Find position of string in another string from end

Above we searched from the beginning of the string. We can also search from end using [`strings.LastIndex`](https://golang.org/pkg/strings/#LastIndex):

```go
s := "hello and second hello"
toFind := "hello"
idx := strings.LastIndex(s, toFind)
fmt.Printf("when searching from end, '%s' is in s at position %d\n", toFind, idx)
```

## Find all occurences of a substring

Above we only found first occurence of substring. Here's how to find all of them:
```go
s := "first is, second is, third is"
toFind := "is"
currStart := 0
for {
      idx := strings.Index(s, toFind)
      if idx == -1 {
            break
      }
      fmt.Printf("found '%s' at position %d\n", currStart + idx)
      currStart += idx + len(toFind)
      s = s[idx + len(toFind):]
}
```

<!-- TODO: example using a regular expression -->

## Check if a string contains another string

Using [`strings.Contains`](https://golang.org/pkg/strings/#Contains):

```go
s := "is hello there?"
toFind := "hello"
if strings.Contains(s, toFind) {
      fmt.Printf("'%s' contains '%s'\n", s, toFind)
} else {
      fmt.Printf("'%s' doesn't contain '%s'\n", s, toFind)
}
```

## Check if a strings starts with another strings

Using [`strings.HasPrefix`](https://golang.org/pkg/strings/#HasPrefix):

```go
s := "this is string"
toFind := "this"
if strings.HasPrefix(s, toFind) {
      fmt.Printf("'%s' starts with '%s'\n", s, toFind)
} else {
      fmt.Printf("'%s' doesn't start with '%s'\n", s, toFind)
}
```

## Check if a strings ends with another strings

Using [`strings.HasSuffix`](https://golang.org/pkg/strings/#HasSuffix):

```go
s := "this is string"
toFind := "string"
if strings.HasPSuffix(s, toFind) {
      fmt.Printf("'%s' ends with '%s'\n", s, toFind)
} else {
      fmt.Printf("'%s' doesn't end with '%s'\n", s, toFind)
}
```
