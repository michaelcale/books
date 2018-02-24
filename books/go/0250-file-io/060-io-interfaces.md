---
Title: I/O related interfaces
Id: 801000b6
---
Go standard library defines several [interfaces](ch-1221) related to i/o.

They are crucial for abstracting i/o operation from the concrete object.

Thanks to `io.Reader` interface we can write code that operates on any type that implements that interface, be it type representing a file on disk, a network connection or a buffer in memory.

Having, for example, a JSON decoder operate on `io.Reader` is more powerful than JSON decoder that can only work on files.

For maximum flexibility, when possible you should write functions that operate on interfaces like `io.Reader` or `io.Writer` and not concrete types like `*os.File`.

## `io.Reader`

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

`io.Reader` is a crucial abstraction for reading from sequential stream of bytes.

`Read` function reads up to `len(p)` bytes int a buffer `p`.

It might return less than the `len(p)`, even 0 bytes.

When it reaches end of file, it returns `io.EOF` as error. Notice that `Read` is allowed to return data in the same call as returning `io.EOF` so handling end of file requires attention to details.

`io.Reader` doesn't allow going back in the stream. For that the type must implement `io.Seeker` or `io.ReaderAt` interfaces.

## `io.Writer`

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

`io.Writer` is for writing to a sequential streawm of bytes.

`Write` writes bytes in `p` and returns number of written bytes and an error.

`Write` guarantees that it'll write all data or return an error i.e. if returned n is < len(p) then err must be non-nil.

## `io.Closer`

```go
type Closer interface {
        Close() error
}
```

`Closer` describes streams that msut be explicitly closed.

`Close` returns an error because it's required in some real-world cases. For example, when doing buffered writes to a file, `Close` might need to flush remaining buffered data to a file, which might fail.

For that reason it's important to check error returned from `Close` when closing a write-able streams.

## `io.ReaderAt`

```go
type ReaderAt interface {
        ReadAt(p []byte, off int64) (n int, err error)
}
```

`io.ReaderAt` is like `io.Reader` but allows to read at any position in the stream.

This is possible in files but not in network connections.

## `io.WriterAT`

```go
type WriterAt interface {
        WriteAt(p []byte, off int64) (n int, err error)
}
```

`io.WriterAt` is like `io.Write` but allows to write at an arbitrary position in the stream.

## `io.Seeker`

```go
type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
}
```

`io.Seeker` allows seeking within the stream. If you can seek, you can also implement `io.ReaderAt` and `io.WriterAt`.

