---
Title: Non-blocking receive with select
Id: 80100096
---
You can use `default` part of `select` statement to do a non-blocking wait.

@file non_blocking_wait.go output sha1:a6f654f50355fa2d2f23b635c97b6b042d0772e0 goplayground:DZjC_jJXuWj

During first iteration of `for` loop `select` immediately ends up in `default` clause because channel is empty.

We send a value to the channel there so the next select will pick up the value from the channel.
