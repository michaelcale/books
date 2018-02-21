---
Title: Package naming best practices
Id: 6680
---

It's best if package name is the same as the name of last path element of import path.

If your GitHub account is `github.com/kjk` and you're writing package `markdown`, it's import path should be `github.com/kjk/markdown`.

A common mistake is to include `go` as part of import path e.g. naming the package `markdown` and giving it import path `github.com/kjk/go-markdown`.
