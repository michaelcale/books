---
Title: Decoding a String
Id: 15707
Score: 0
---

```go
decoded, err := base64.StdEncoding.DecodeString(`biws`)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%s", decoded)
// Output: n,,
```

[Playground](https://play.golang.org/p/h2qngYncRs)
