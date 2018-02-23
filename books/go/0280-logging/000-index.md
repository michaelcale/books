---
Title: Logging
Id: 3724
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

@file log.go output

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

@file log_file.go

## Logging to

When running on Unix, we might log to syslog:

@file log_syslog.go
