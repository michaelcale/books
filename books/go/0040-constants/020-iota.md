---
Title: iota
Id: 80100036
---
`iota` makes it easy to declare number constants whose values grow.

Basics:

@file iota.go output

`iota` sets the value of `Low` to 0 and instructs the compiler following constants have increasing numeric values.


## Creating bitmask values with `iota`

Iota can be very useful when creating a bitmask. For instance, to represent the state of a network connection which may be secure, authenticated, and/or ready, we might create a bitmask like the following:

@file iota_2.go output

## Skipping values

The value of `iota` is still incremented for every entry in a constant list even if iota is not used:

@file iota_3.go output

it will also be incremented even if no constant is created at all, meaning the empty identifier can be used to skip values entirely:

@file iota_4.go output

## Using `iota` in an expression list

Because `iota` is incremented after each [`ConstSpec`](https://golang.org/ref/spec#ConstSpec), values within the same expression list will have the same value for `iota`:

@file iota_5.go output

## Using `iota` in an expression

`iota` can be used in expressions, so it can also be used to assign values other than simple incrementing integers starting from zero. To create constants for SI units, use this example from [Effective Go](https://golang.org/doc/effective_go.html#initialization):

@file iota_6.go output
