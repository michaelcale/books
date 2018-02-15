---
Title: Signalling channel with chan struct{}
Id: 801000u5
---
Sometimes you don't want to send a value over a channel but use it only as a way to signal an event.

In those cases use `chan struct{}` to document the fact that the value sent over a channel has no meaning.

Signalling channel is often used as a way to tell goroutines to finish.

@file signalling.go output
