---
Title: Raw access to command line arguments
Id: 172
SOId: 14047
---
To parse command-line arguments, you can use package [`flag`](https://golang.org/pkg/flag/) in stndard library, use a library like [cobra](173).

You can also parse the arguments yourself.

@file raw_cmd_line_args.go output noplayground

The above output is a result of `go run $file -echo echo-arg additional arg`.

## Raw arguments

Raw command-line arguments can be accessed via `[]string` slice `os.Args`.

First element is name of the executable.

Remaining elements are cmd-line arguments. Shell quoting rules apply.

On Windows natively cmd-line arguments are a single UTF-16 Unicode string. Go runtime converts them to UTF-8 string and splits into elements to unify handling across different operating systems.
