---
Title: Writing files
Id: 3332
---

## Write the whole file

The simplest way to write data to a file:

```go
d := []byte("content of the file")
err := ioutil.WriteFile("foo.txt", d, 0644)
if err != nil {
    log.Fatalf("ioutil.WriteFile failed with '%s'\n", err)
}
```

## Open file for writing

```go
f, err := os.Create("foo.txt")
if err != nil {
    log.Fatalf("os.Open failed with '%s'\n", err)
}
```

`Create` returns `*os.File` which implements `io.Writer` and `io.Closer` interfaces.

If file doesn't exist, it'll be created.

If file does exist, it'll be truncated.

You should always close files to avoid leaking file descriptors.

Be aware that `Close` on a file can return an error so for a robust code you should check for errors from `Close`.

Writes can be buffered and `Close` might need to flush remaining cached bytes to disk. That might fail and return an error.

## Open file for appending

TODO: write me

## Write to a file

```go
d := []byte("data to write")
nWritten, err := f.Write(d)
```

