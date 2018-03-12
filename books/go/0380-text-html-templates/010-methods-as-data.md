---
Title: Methods as data
Id: 338
---

In a template `{{ .Foo }}` will access either a field `Foo` in a struct or will call a function `Foo()`:

@file methods_as_data.go output sha1:eb52f9493001b06bed701f265aaf2901ce62f59a goplayground:SJwK4yfVwiY
