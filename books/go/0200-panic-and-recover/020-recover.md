---
Title: Recover
Id: 17010
---
Recover attempts to recover from a `panic`.

The recover *must* be attempted in a deferred statement as normal execution flow has been halted.

The `recover` statement must appear *directly* within the deferred function enclosure.

Recover statements in functions called by deferred function calls will not be honored.

The `recover()` call will return the argument provided to the initial panic, if the program is currently panicking.

If the program is not currently panicking, `recover()` will return `nil`.

@file recover.go output allow_error
