---
Title: Handling errors in short programs
Id: 125
SOId: k1k100qf
---
Propagating errors up to the callers is tedious. You end up with many lines looking like:

```go
r, err := os.Open("my file")
if err != nil {
    return err
}
```

This kind of error handling diligence is crucial for writing robust software.

Sometimes you write shorter cmd-line programs where such discipline is not warranted.

You can simplify error handling with a helper functions `PanicIfErr(err error, args ...interface{})`:

```go
r, err := os.Open("my file")
PanicIfErr(err)
```

Implementation of such helper:

@file errors_in_short_programs.go output allow_error sha1:b71b4c3a3d33c20785aa36fd6d66156649a8b5d8 goplayground:v2lHFqenkOG