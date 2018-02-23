---
Title: I/O related interfaces
Id: 801000b6
---
Go standard library defines several [interfaces](a-1221) related to i/o.

They are crucial for abstracting i/o operation from the concrete object.

Thanks to `io.Reader` interface we can write code that operates on any type that implements that interface, be it type representing a file on disk, a network connection or a buffer in memory.

Having, for example, a JSON decoder operate on `io.Reader` is more powerful than JSON decoder that can only work on files.

## `io.Reader`

TODO: write me

## `io.Writer`

TODO: write me

## `io.Closer`

TODO: write me

## `io.ReaderAt`

TODO: write me

## `io.Seeker`

TODO: write me

