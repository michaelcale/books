---
Title: Portable file path operations
Id: rd6000ld
---
Unfortunately different operating systems have different rules about the format of file paths.

For example, on Unix and Mac OS, path separator character is `/` and on Windows it's `\`.

For portable programs, it's important to use functions in `filepath` package that understand conventions used by a given operating system.

Note: `filepath` package manages OS file paths. There's also `path` package with similar functionality but it always uses `/` as path separator.

## Join a path

```go
path := filpath.Join("dir", "file.txt")
```

You can join more than 2 path elements.

On Windows the above would return `dir\file.txt`, on Unix and Mac OS it would return `dir/file.txt`.

## Split a path into a directory and file

```go
path := filepath.Join("dir", "file.txt")
dir, file := filepath.Split(path)
```

## Split path into all components

```go
path := filepath.Join("dir", "subdir", "file.txt")
parts := filepath.SplitList(path)
// parts = []string{"dir", "subdir", "file.txt"}
```

## Get file name from path

```go
path := filepath.Join("dir", "file.txt")
file := filepath.Base(path)
// file = "file.txt"
```

## Get directory name from path

```go
path := filepath.Join("dir", "file.txt")
dir := filepath.Dir(path)
// dir = "dir"
```

## Get file extension

```go
ext := filepath.Ext("file.txt")
// ext = ".txt"
```

