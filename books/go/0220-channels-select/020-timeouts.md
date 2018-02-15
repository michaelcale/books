---
Title: Timout reading from channel with select
Id: 6050
---
Receiving froma channel with `<- chan` or `for range` loop blocks.

Sometimes you want to limit time waiting for a value on a channel.

It's possible with `select`:

@file timeout.go output
