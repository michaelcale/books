---
Title: Executing commands
Search: Start process, execute process
Id: 1097
---
Package `exec` in standard library is a cross-platform way to launch processes, capture their output and more.

## Basic command execution

The simplest usage is:
* create `exec.Cmd` struct using `exec.Command(exe string, args ...string)`
* call `cmd.CombinedOutput()` to execute the cmd and get combined stdout and stderr
* to get only stdout, call `cmd.Output()`

@file index.go output noplayground

## More advanced command execution

@file index2.go output noplayground

This is functionally the same as the above example but we use a more fine-grained control.

We capture stdout and stderr of the process by setting `cmd.Stdout` and `cmd.Stderr` to a memory-backed `io.Writer`.

`cmd.Start()`starts command as new OS process. It executes concurrently with our code, as OS processes do.

We need to call `cmd.Wait()` to wait for the process to finish. To prevent waiting infinitely, you might want to add a [timeout](a-3521).