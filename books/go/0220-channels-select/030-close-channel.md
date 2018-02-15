---
Title: Closing channels
Id: rd6000v9
---
You can [close](a-rd6000v9) a channel with `close(chan)`.

Main purpose of closing a channel is to notify worker goroutines that their work is done and they can finish.

This ensures you don't leak goroutines.

**Closing a channel ends `range` loop over it:**

@file close.go output

**A receive from closed channels returns zero value immediately:**

@file close2.go output

**When receiving, you can optionally test if channel is closed:**

@file close3.go output

**Closing channel twice [panics](ch-4350):**

@file close4.go output allow_error

**Sending to closed channel [panics](ch-4350).**

@file close5.go output allow_error
