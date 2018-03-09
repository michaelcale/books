---
Title: Cross compilation by using gox
Id: 271
Score: 2
SOId: 30867
---
Another convenient solution for cross compilation is the usage of `gox`: https://github.com/mitchellh/gox

## Installation
The installation is done very easily by executing `go get github.com/mitchellh/gox`. The resulting executable gets placed at Go's binary directory, e.g. `/golang/bin` or `~/golang/bin`. Ensure that this folder is part of your path in order to use the `gox` command from an arbitrary location.

## Usage
From within a Go project's root folder (where you perform e.g. `go build`), execute `gox` in order to build all possible binaries for any architecture (e.g. x86, ARM) and operating system (e.g. Linux, macOS, Windows) which is available.

In order to build for a certain operating system, use e.g. `gox -os="linux"` instead. Also the architecture option could be defined: `gox -osarch="linux/amd64"`.
