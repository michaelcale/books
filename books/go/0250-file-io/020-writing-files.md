---
Title: Writing files
Id: 161
SOId: 3332
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

```go
f, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
if err != nil {
    log.Fatalf("os.Open failed with '%s'\n", err)
}
```

Second argument to `os.OpenFile` is a flag that determines the exact mode for opening the file.

When you open for reading, use `os.Open`.

When you open for writing to a new file, use `os.Create`.

When you open for appending to existing file, use `os.OpenFile` with the following flags:
* `os.O_WRONLY`. Could also be `os.RDWR` if we also to both read and write
* `os.O_APPEND` means that if file exists, we'll append
* `os.O_CREATE` means that if file doesn't exist, we'll create it. Without this flag opening non-existing file would fail

## Write to a file

```go
d := []byte("data to write")
nWritten, err := f.Write(d)
```

## File permissions when creating files

When you create a new file with `os.Create` or `os.OpenFile`, you need to provide file permissions for the new file.

For most cases `0644` is a good choice.

Those are Unix style permissions in octal format.

Let's deconstruct parts of `0644`
* 0 means this is a number in octal format. That means that each number is in 0-7 range (compare to regular decimal notation where each number is in 0-9 range)
* 6 is permissions for the creator of the file. It's a bitmask 4 + 2. 4 means read permissions, 2 means write permission. When combined it means read/write access
* 4 is permissions for the group that user belongs to. I.e. every user that also belongs to the same group will have those permissions. 4 means only read permissions
* last 4 is for everyone. Again, 4 means only read permissions
* for completness, 1 represents exuectuable permissions so if this is meant to be an executable file, the permissions flag would be `0755` i.e. for each user/group/anonymous component we add 1 to set executable bit

Those permissions are native to Unix and Mac OS but they map loosely on Windows.

To manage Windows permissions you would need to use Windows API.
