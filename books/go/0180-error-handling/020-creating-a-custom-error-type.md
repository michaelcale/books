---
Title: Create custom error type
Id: 2706
---

In Go any type that implements built-in `error` interface (i.e.
`Error() string` method) can be used as error:

@file create_custom_error.go output sha1:89a9cf70397bd90ebd38e6c16761c79753fdc67b goplayground:LQ_lhw9zaJP
