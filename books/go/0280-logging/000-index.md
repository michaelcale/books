---
Title: Logging
Id: 174
SOId: 3724
---
Logging is a very deep subject because different programs have different logging requirements.

## Loging with `fmt.Printf` and `fmt.Fprintf`

The simplest way to log is to write to standard output (stdout):

```go
fmt.Printf("Logging to %s\n", "stdout")
```

To write to standard error (stderr):
```
fmt.Fprintf(os.Stderr, "Logging to %s\n", "stderr")
```

## Logging with `log` package

Standard package [`log`](https://golang.org/pkg/log/) offers more functionality:

@file log.go output sha1:7ec653b436ede6b3ad152a355b3e35c556119cf9 goplayground:T6IHD0uxEsC

Compared to `fmt.Printf`, `log.Printf`:
* by default logs to stderr (`os.Stderr`)
* adds current time to each log line
* ensures that echo log is on it's own line by adding `\n` if not explicitly provided

To log fatal issues:
```go
f, err := os.Open("file.txt")
if err != nil {
    log.Fatalf("os.Open('file.txt') failed with '%s'\n", err)
}
```

`log.Fatalf` logs the message and calls `os.Exit(1)` to end the process.

## Logging to a file

Log package allows changing where the log output is sent. We can log to a file:

@file log_file.go sha1:0dd5e3acabedd5087ef46852fb9848cb76a8bb71 goplayground:thVsmxVof2c

## Logging to

When running on Unix, we might log to syslog:

@file log_syslog.go sha1:f2aca065f394259d2a04e0ceab84e363493112dd goplayground:qENpElxIRMS
