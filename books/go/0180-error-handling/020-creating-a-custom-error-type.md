---
Title: Create custom error type
Id: 2706
---

In Go any type that implements built-in `error` interface (i.e.
`Error() string` method) can be used as error:

@file create_custom_error.go output
