---
Title: Reading files
Id: 801000of
---
## Read the whole file

The simplest way to read a whole file is:

```go
d, err := ioutil.ReadFile("foo.txt")
if err != nil {
    log.Fatalf("ioutil.ReadFile failed with '%s'\n", err)
}
fmt.Printf("Size of 'foo.txt': %d bytes\n", len(d))
```

## Open file for reading, close file

```go
f, err := os.Open("foo.txt")
if err != nil {
    log.Fatalf("os.Open failed with '%s'\n", err)
}
defer f.Close()
```

Open returns `*os.File` which implements `io.Reader` and `io.Closer` interfaces.

You should always close files to avoid leaking file descriptors. `defer` is perfect for ensuring `Close` will be called on exit of the function.

## Read file line by line

@file read_lines.go output noplayground

