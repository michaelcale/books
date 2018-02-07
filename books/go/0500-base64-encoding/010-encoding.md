---
Title: Encoding
Id: 15704
Score: 0
---
```
const foobar = `foo bar`
encoding := base64.StdEncoding
encodedFooBar := make([]byte, encoding.EncodedLen(len(foobar)))
encoding.Encode(encodedFooBar, []byte(foobar))
fmt.Printf("%s", encodedFooBar)
// Output: Zm9vIGJhcg==
```
[Playground](https://play.golang.org/p/A5c_BSMFrQ)
