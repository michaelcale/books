---
Title: Closing channels
Id: rd6000v9
---

TODO: write me

You can [close](a-rd6000v9) a channel with `close(chan)`.

Closing channel twice [panics](ch-4350).

Sending to closed channel [panics](ch-4350).

A receive from closed channels returns zero value immediately.

A closed channel finishes `range` over a channel.
