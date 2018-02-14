---
Title: Error Handling
Id: 785
---
Basic error handling:

@file index.go output

**Important notes about error handling:**

Unlike languages like C# or Python, Go handles errors by returning error values, not raising exceptions.

Go also includes exception handling with [panic and recover](ch-4350) but it's not supposed to be used frequently.

Errors are values, just like integers or string.

Type `error` is a built-in [interface](ch-1221) which implements `Error() string` method.

You can use your own types as `error` values by implementing `Error() string` method or you can use `errors.New(msg string)` or `fmt.Errorf(format string, args... interface{})` from standard library.

To indicate no errors, return `nil`.

If a function returns errors, it should always be the last returned value.

Often you want to propagate error value up to the caller until a point in your code where you want to handle it.

**You should always check for errors.**

You can often see people asking on mailing lists or StackOverflow why their code doesn't and it would be obvious if the code didn't skip error checking.

