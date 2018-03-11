---
Title: range over a channel
Id: 110
SOId: 801000o5
---

Iterating over a channel with `range`:

@file range_channel.go output

`range` over a channel receives elements sent over a channel.

Iteration ends when channel is closed.

This is a more convenient version of:

@file range_channel2.go output
