---
Title: iota
Id: 80100036
---
`iota` makes it easy to declare number constants whose values grow.

Basics:

@file iota.go output sha1:a8b1274df4bb37c5d29db536bbf026e89f600e17 goplayground:Hom7xSeu2_V

`iota` sets the value of `Low` to 0 and instructs the compiler that the following constants have increasing numeric values.


## Creating bitmask values with `iota`

Iota can be very useful when creating a bitmask. For instance, to represent the state of a network connection which may be secure, authenticated, and/or ready, we might create a bitmask like the following:

@file iota_2.go output sha1:6d3b494ed9a7f5412cecb9934a3f107b27511419 goplayground:PgPiSJeXXtP

## Skipping values

The value of `iota` is still incremented for every entry in a constant list even if iota is not used:

@file iota_3.go output sha1:85d09c78773f11dadfcb5b117979571188019251 goplayground:QXKxVfUW9e-

It will also be incremented even if no constant is created at all, meaning the empty identifier can be used to skip values entirely:

@file iota_4.go output sha1:73888bb42c8811dfa670abcb058a736e93999a9e goplayground:60ov9hRLZK1

## Using `iota` in an expression list

Because `iota` is incremented after each [`ConstSpec`](https://golang.org/ref/spec#ConstSpec), values within the same expression list will have the same value for `iota`:

@file iota_5.go output sha1:fb4d2417b5cca2718b8e4d1e4f4faeab57b3df92 goplayground:ft9XGtyS3ur

## Using `iota` in an expression

`iota` can be used in expressions, so it can also be used to assign values other than simple incrementing integers starting from zero. To create constants for SI units, use this example from [Effective Go](https://golang.org/doc/effective_go.html#initialization):

@file iota_6.go output sha1:20fbf7ae725e839722cff1187c46a7b1316c9da9 goplayground:QpLaWO0cLyP
