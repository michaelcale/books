---
Title: Methods as data
Id: 338
---

In a template `{{ .Foo }}` will either access struct field `Foo` or call a function `Foo()`:

@file methods_as_data.go output sha1:eb52f9493001b06bed701f265aaf2901ce62f59a goplayground:SJwK4yfVwiY
