---
Title: Directory operations
Id: 163
SOId: 9310
---
## Create directory

```go
dir := "my_dir"
err := os.Mkdir(dir, 0755)
if err != nil {
    fmt.Printf("os.Mkdir('%s') failed with '%s'\n", dir)
}
```

```go
dir := filepath.Join("topdir", "subdir")
err := os.MkdirAll(dir, 0755)
if err != nil {
    fmt.Printf("os.MkdirAll('%s') failed with '%s'\n", dir)
}
```

`os.Mkdir` only succeeds if parent directory of dir already exists.

`os.MkdirAll` will create all intermediary directories.

`0755` describes permissions for the directory.

Those are Unix style permissions in octal format.

Let's deconstruct parts of `0755`
* 0 means this is a number in octal format. That means that each number is in 0-7 range (compare to regular decimal notation where each number is in 0-9 range)
* 7 is permissions for the creator of the file. It's a bitmask 4 + 2 + 1. 4 means read permissions, 2 means write permission, 1 means that directory can be traversed. When combined it means read/write/traverse access
* 5 is permissions for the group that user belongs to. I.e. every user that also belongs to the same group will have those permissions. 4 + 1 means read and traverse permissions, but not write permissions
* last 5 is for everyone. Again, 5 means read and traverse permissions

## Delete directory

```go
dir := "my_dir"
err := os.Remove(dir)
if err != nil {
    fmt.Printf("os.Remove('%s') failed with '%s'\n", path, err)
}
```

`os.Remove` only works for empty directories i.e. directories without any files of sub-directories.

```go
dir := "my_dir"
err := os.RemoveAll(dir)
if err != nil {
    fmt.Printf("os.RemoveAll('%s') failed with '%s'\n", path, err)
}
```

`os.RemoveAll` removes the directory and all its children (files and sub-directories).

## List files in a directory

To list files and directories in a given directory we can use `ioutil.ReadDir`:

@file dir_list_files.go output noplayground

## List files recursively

@file filepath_walk.go output noplayground

To visit files in a directory recursively use `filepath.Walk`.

You provide a callback function that will be called for every file and directory under directory.

Callback function is called even for files/directories that we can't read e.g. due to insufficient permissions.

We can end traversal early by returning non-nil error from callback function.