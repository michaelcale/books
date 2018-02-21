---
Title: Creating a package
Id: 801000h5
---
So you want to create a Go package.

You can use any source code hosting service. For this tutorial let's assume that you're using https://github.com, your account is `kjk` and you want to create package `foo`.

Create a repository `foo` on GitHub.

Create a directory `$GOPATH/src/github.com/kjk` and checkout repository `foo` there:
```bash
$ mkdir -p $GOPATH/src/github.com/kjk
$ cd $GOPATH/src/github.com/kjk
$ git clone github.com/kjk/foo
```

Create a file `main.go`:
```go
package foo

func Bar() string {
    return "bar"
}
```

This is a simplest package that defines function `Bar`.

Notice that package name `foo` matches directory name `foo`. It's not required  but recommended.
