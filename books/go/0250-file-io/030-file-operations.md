---
Title: File operations
Id: 3333
---
## Get file size

@file file_size.go output noplayground

Instead of `os.Stat` we can also use `os.Lstat`. The difference is that `os.Stat` follows symbolic links and `os.Lstat` doesn't.

In other words: for a symbolic link `os.Lstat` returns information about link and `os.Stat` about the file that it links to.

## Get information about the file

@file file_info.go output noplayground

## Check if a file exists

@file file_exists.go output noplayground

Checking if a file exists is surprisingly tricky and it's impossible to write a generic function that handles all the nuances.

Here are decisions we made:
* it treats files and directories the same. If a path exists and you want to distinguish between directory and a file, you would need to call `IsDir()` on result of `os.Lstat`
* if a file is a symbolic link, do test the link or the real file it links to? We used `os.Lstat` so we test the link. We could also use `os.Stat` to resolve the symbolic link
* "path doesn't exist" is only one of possible errors returned by `os.Lstat`. Do we want to distinguish between "file doesn't exist" and "file exists and we don't have access"? We decided to be more informative but in some cases it would be simpler to just return a bool and always return false if `os.Lstat` fails

## Delete a file

```go
path := "foo.txt"
err := os.Remove(path)
if err != nil {
    if os.IsNotExist(err) {
        fmt.Printf("os.Remove failed because file doesn't exist\n")
    } else {
        fmt.Printf("os.Remove failed with '%s'\n", err)
    }
}
```

`os.Remove` returns an error for files that don't exist.

Usually you want to ignore such errors which you can do by testing error with `os.IsNotExist(err)`.

## Rename a file

```go
oldPath := "old_name.txt"
newPath := "new_name.txt"
err := os.Rename(oldPath, newPath)
if err != nil {
    fmt.Printf("os.Rename failed with '%s'\n", err)
}
```

## Copy a file

@file file_copy.go sha1:2f3f1dbd0d3c95b473802cbc841610f161f8193d goplayground:Mr3njneW17O

Writing a generic function for copying files is tricky and it's impossible to write a function that serves all use cases.

Here are policy decisions we made:
* should we over-write existing files or return error if destination exists? We decided to over-write
* what permissions should new file have? We decided for the simplest case of using default permissions. Another option would be to copy permissions from source or allow the caller to provide permissions

If you want different behavior, you will have to modify the code as needed.
