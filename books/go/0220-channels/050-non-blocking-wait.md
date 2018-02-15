---
Title: Non-blocking receive
Id: 80100096
---
You can use `default` part of `select` statement to do a non-blocking wait.

@file non_blocking_wait.go output

During first iteration of `for` loop `select` immediately ends up in `default` clause because channel is empty.

We send a value to the channel there so the next select will pick up the value from the channel.
