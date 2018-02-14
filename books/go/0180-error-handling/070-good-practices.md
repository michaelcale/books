---
Title: Good practices for error messages
Id: k1k100a6
---

Good error messages help in debugging problems.

In production deployments error messages are logged. When something goes wrong log file is the main source for debugging clues.

When creating error messages, think about that and include as much useful information as possible.

Bad error message:

```go
func openLog(path string) (io.Reader, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, errors.New("failed to open log file")
    }
    ...
}
```

Better error message:

```go
func openLog(path string) (io.Reader, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, errors.New("openLog('%s'), os.Open() failed with '%s'", path, err)
    }
    ...
}
```

Second version of error message includes more debugging clues:
* which function generated the error
* path of the file we were trying to open
* the error message from os.Open() so that we can tell if it was because the file doesn't exist or we don't have access rights to it etc.
