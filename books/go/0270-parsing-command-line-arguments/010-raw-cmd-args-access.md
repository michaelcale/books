---
Title: Raw access to command line arguments
Id: 172
SOId: 14047
---

If `flag` package or a third-party library doesn't provide the features you want, you can parse the arguments yourself.

@file raw_cmd_line_args.go output noplayground

The above output is a result of `go run $file -echo echo-arg additional arg`.

## Raw arguments

Raw command-line arguments can be accessed via `[]string` slice `os.Args`.

First element is name of the executable.

Remaining elements are cmd-line arguments as decoded by OS shell.

On Windows cmd-line arguments are a single UTF-16 Unicode string.

Go runtime converts them to UTF-8 string and splits into separate arguments to unify handling across different operating systems.

