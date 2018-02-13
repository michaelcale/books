---
Title: CSV
Id: 20508
---
Consider this CSV data:

```csv
#id,title,text
1,hello world,"This is a ""blog""."
2,second time,"My
second
entry."
```

This data can be read with the following code:

```go
// r can be any io.Reader, including a file.
csvReader := csv.NewReader(r)
// Set comment character to '#'.
csvReader.Comment = '#'
for {
    post, err := csvReader.Read()
    if err != nil {
        log.Println(err)
        // Will break on EOF.
        break
    }
    fmt.Printf("post with id %s is titled %q: %q\n", post[0], post[1], post[2])
}
```

And produce:

```text
post with id 1 is titled "hello world": "This is a \"blog\"."
post with id 2 is titled "second time": "My\nsecond\nentry."
2009/11/10 23:00:00 EOF
```

[Playground](https://play.golang.org/p/d2E6-CGGIe).
