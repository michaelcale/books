---
Title: Format text
Id: 29829
---

Go's standard library implements C-style string formatting in the [`fmt`](https://golang.org/pkg/fmt/) package.

@file formatting_text.go output sha1:8320ec396f0bcaa1ef717aad667081c2df7991ea goplayground:j5SLVfc6RdI

The first argument to `fmt.Sprintf` is a formatting string which defines how to format subsequent arguments. Subsequent arguments are the values that will be formatted.

`fmt.Sprintf` creates a formatted string.

For convenience, there's also:
* `fmt.Fprintf(w io.Writer, string format, args... interface{})`, which will write a formatted string to a given writer
* `fmt.Pritnf(format string, args.. interface{})` which writes a formatted string to `os.Stdout`.

<!-- TODO: more examples for goal-oriented -->

The function `Sprintf` formats the string in the first parameter, replacing the verbs with the values of the subsequent parameters and returns the result. Like `Sprintf`, the function `Printf` also formats but instead of returning the result it prints the string.

## List of string formatting verbs

```text
%v    // the value in a default format
        // when printing structs, the plus flag (%+v) adds field names
%#v   // a Go-syntax representation of the value
%T    // a Go-syntax representation of the type of the value
%%    // a literal percent sign; consumes no value
```

Boolean:

```text
%t    // the word true or false
```

Integer:

```text
%b    // base 2
%c    // the character represented by the corresponding Unicode code point
%d    // base 10
%o    // base 8
%q    // a single-quoted character literal safely escaped with Go syntax.
%x    // base 16, with lower-case letters for a-f
%X    // base 16, with upper-case letters for A-F
%U    // Unicode format: U+1234; same as "U+%04X"
```

Floating-point and complex constituents:

```text
%b    // decimalless scientific notation with exponent a power of two,
        // in the manner of strconv.FormatFloat with the 'b' format,
        // e.g. -123456p-78
%e    // scientific notation, e.g. -1.234456e+78
%E    // scientific notation, e.g. -1.234456E+78
%f    // decimal point but no exponent, e.g. 123.456
%F    // synonym for %f
%g    // %e for large exponents, %f otherwise
%G    // %E for large exponents, %F otherwise
```

String and slice of bytes (treated equivalently with these verbs):

```text
%s    // the uninterpreted bytes of the string or slice
%q    // a double-quoted string safely escaped with Go syntax
%x    // base 16, lower-case, two characters per byte
%X    // base 16, upper-case, two characters per byte
```

Pointer:

```text
%p    // base 16 notation, with leading 0x
```
