---
Title: Closing channels
Id: rd6000v9
---
You can [close](a-rd6000v9) a channel with `close(chan)`.

Main purpose of closing a channel is to notify worker goroutines that their work is done and they can finish.

This ensures you don't leak goroutines.

**Closing a channel ends `range` loop over it:**

@file close.go output sha1:82f3912d3a0d828d24d2b3f41cb8089a44f559f2 goplayground:q3mX8S36CBK

**A receive from closed channels returns zero value immediately:**

@file close2.go output sha1:489b17f8a7bd30ff4e98092919951c20f4de9e54 goplayground:a16qGMRfQwq

**When receiving, you can optionally test if channel is closed:**

@file close3.go output sha1:2f2778e77c917d68a9a6ae0a7f51904cfb0634dd goplayground:CtXiCamXQFr

**Closing channel twice [panics](ch-4350):**

@file close4.go output allow_error sha1:d57ed196ecf819fe95ebe0aeb5557c97a115ce2a goplayground:ApSna0RUYjp

**Sending to closed channel [panics](ch-4350).**

@file close5.go output allow_error sha1:fa33388968966773ecdccb154c0a2b8bd94dad96 goplayground:D9p7YgUVKM1
