Title: Setting up the environment
Id: 14406
Score: 61
Body:
If Go is not pre-installed in your system you can go to https://golang.org/dl/ and choose your platform to download and install Go.

To set up a basic Go development environment, only a few of the many environment variables that affect the behavior of the `go` tool (See: [Listing Go Environment Variables][1] for a full list) need to be set (generally in your shell's `~/.profile` file, or equivalent on Unix-like OSs).

## `GOPATH`

Like the system `PATH` environment variable, Go path is a `:`(`;` on Windows) delimited list of directories where Go will look for packages. The `go get` tool will also download packages to the first directory in this list.

The `GOPATH` is where Go will setup associated `bin`, `pkg`, and `src` folders needed for the workspace:

 - `src` — location of source files: `.go`, `.c`, `.g`, `.s`
 - `pkg` — has compiled `.a` files
 - `bin` — contains executable files built by Go

From Go 1.8 onwards, the `GOPATH` environment variable will have a [default value][2] if it is unset. It defaults to $HOME/go on Unix/Linux and %USERPROFILE%/go on Windows.

Some tools assume that `GOPATH` will contain a single directory.

## `GOBIN`

The bin directory where `go install` and `go get` will place binaries after building `main` packages. Generally this is set to somewhere on the system `PATH` so that installed binaries can be run and discovered easily.

## `GOROOT`

This is the location of your Go installation. It is used to find the standard libraries. It is very rare to have to set this variable as Go embeds the build path into the toolchain. Setting `GOROOT` is needed if the installation directory differs from the build directory (or the value set when building).


  [1]: http://stackoverflow.com/documentation/go/198/introduction-to-go/14405/listing-go-environment-variables#t=20160805180807745607
  [2]: https://golang.org/doc/go1.8#gopath
|======|
