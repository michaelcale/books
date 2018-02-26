---
Title: Timing out process execution
Id: 3521
---

There's no guarantee that a process will finish running so you might want to add a timeout.

@file time_out.go output

Before calling `cmd.Wait()` we use `time.AfterFunc()` to kill the process after a timeout expired.

`cmd.Process.Kill()` makes `cmd.Wait()` exit with an error (`sinal: killed` on Unix).

If `cmd.Wait()` finishes before a timeout, we cancel the timer.

Variable `timedOut` is `int32` and not `bool`.

We use `atomic.StoreInt32` and `atomic.LoadInt32` becase there is a subtle timing issues.

If a process finishes as usual, after `cmd.Wait()` exits and before we call `stopTimer.Stop()`, the timer might expire and execute the function anyway.

This is exceedingly unliley but something to keep in mind.

