---
Title: Pull (streaming) XML parsing
Id: 190
---

Parsing into a struct is convenient but requires a lot of memory to hold the whole decoded document in memory.

In some cases XML files are so large that it's not possible to decode the whole file into memory. For example XML dumps of Wikipedia content are several gigabytes in size.

Pull parsing is more efficient but API is harder to use.

@file pull_parse.go output sha1:9bfb2d72c0d60f67b1a937590d407411844dab8b goplayground:69qT0qN7Fgg

Pull parsing requests next token from stream of XML tokens.

For start tag like `<person>` we get `xml.StartElement` token.

For end tag like `</person>` we get `xml.EndElemnt` token.

For data inside the element `<person>data</person>` we get `xml.CharData` token.

When decoder reaches the end, it returns error `io.EOF`.

In the above example we print `age` attribute of `<person>` element and char data inside `<city>` element.

This is a very basic example. In real programs you might need to remember more state.

For example, if your XML is:

```
<foo>
  <bar>
    <foo></foo>
  </bar>
</foo>
```

If you look just at the `xml.StartElement` token, you don't know if `foo` is for the top-level element or is it a child of `<bar>` element.
