---
Title: Installing packages locally with go get
Id: 154
SOId: 80100042
---

In order to import a package in your program you must first download the package locally.

`go get -u github.com/gomarkdown/markdown` downloads package files from https://github.com/gomarkdown/markdown/.

`-u` means `update` i.e. if the package is already downloaded locally, it'll be updated to latest version.

The code will be downloaded into `$GOPATH/src/github.com/gomarkdown/markdown` i.e. package import path is also a name of directory where the package files are stored locally.

Import path is appended to `$GOPATH/src`.

Most of the time import path is also the location of source files but `go get` also allows for redirects via meta-data embedded in HTML files.
