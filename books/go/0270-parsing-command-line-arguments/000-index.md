---
Title: Parse command line arguments
Search: flags
Id: 171
SOId: 4023
---
Go standard library has a [`flag`](https://golang.org/pkg/flag/) package for parsing cmd-line arguments:

@file index.go output noplayground

Output above is a result of calling `go run $file -echo echo-arg additiona arg`.

## Defining arguments

You register cmd-line arguments to be recognized with `flag.IntVar(parsedValue *int, name string, defaultValue int, usage string)`.

There are functions for other common types:
* `flag.BoolVar`
* `flag.DurationVar`
* `flag.Float64Var`
* `flag.IntVar`, `flag.UIntVar`, `flag.Int64Var`, `flag.UInt64Var`
* `flag.StringVar`

If you register string argument with name `echo`, the way to provide it on cmd-line is `-echo ${value}` or `-echo=${value}`.

Command-line argument naming is different from POSIX standard of `--echo` or Windows standard of `/echo`.

For boolean values you can say: `-help` (implicitly true), `-help=true` or `-help=false`.

Saying `-help false` is not what you might expect. Flag `-help` is set to rue and `false` is considered an additional argument.

## Parsing and accessing remaining arguments

After parsing arguments, call `flag.Parse()`.

Parsing can fail if:
* unknown flag was given
* a flag didn't parse based on their type (e.g. it was registred as int but the value was not a valid number)

In case of failure, help text describing flags is shown and program exits wiht error code 2.

You can explicitly pring help text using `flag.Usage()`. This is often triggered by `-help` flag.

Help text is based on usage text provided in `flag.IntVar` and others.

Command-line arguments that don't start with `-` are untouched and can be accessed with `flag.Args()`.

## Limitations

The big features missing from `flag` package:
* no support of POSIX style `--name`, only `-name`
* no support for short alternatives e.g. `-n` being synonym with `--name`

Your options are:
* [raw access to cmd-line arguments](172)
* use a [third party library](173)
