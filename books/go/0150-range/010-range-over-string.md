---
Title: range over a string
Id: 107
SOId: 801000d6
---

## Iterate over bytes

You can iterate over bytes in a string:

@file range_string_bytes.go output

## Iterate over runes

Things are more complicated when you want to iterate over logical characters (runes) in a string:

@file range_string.go output

In Go strings are immutable sequence of bytes. Think a read-only `[]byte` slice.

Each byte is in 0 to 255 range.

There are many more characters in all the world's alphabets.

Unicode standard defines unique value for every known character. Unicode calls them code points and they are integers that can fit in 32 bits.

To represent Unicode code points, Go has a `rune` type. It is an alias for `int32`.

Literal strings in Go source code are UTF-8 encoded.

Every Unicode code point can be encoded with 1 to 4 bytes.

In this form of iteration, Go assumes that a string is UTF-8 encoded. `range` decodes each code point as UTF-8, returns decoded rune and its byte index in string.

You can see the byte index of last code point jumped by 3 because code point before it represents a Chinese character and required 3 bytes in UTF-8 encoding.

<!-- TODO: the same using unicode package -->

<!-- TODO: link to character set conversion article -->

## Aside on strings and UTF-8

Go strings are just slices of bytes. You can put whatever data you want in them.

How the data is interpreted is up to your code.

Go doesn't care if string is a valid UTF-8 sequence. Go doesn't validate if string is as valid UTF-8 sequence.

That being said, Go does include support for working with UTF-8 encoded strings easily.

The behavior of `range` is one example of that.

Go also mandates that Go source code files are valid UTF-8.
