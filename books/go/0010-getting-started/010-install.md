Title: Installing Go compiler
Id: 20381
Body:

There are 2 production quality compilers:
* gc, the official compiler
* [gccgo](https://golang.org/doc/install/gccgo). It's also created by the Go project, but is not used as ofent as gc.

This article describes installing gc compiler.

## Installing on Windows

Download `.msi` installer from [https://golang.org/dl/](https://golang.org/dl/) and run it.

Read [setting up Windows environment](a-rd600086).

## Installing on Mac OS

### Using official binaries

Download `.pkg` installer from [https://golang.org/dl/](https://golang.org/dl/) and run it.

### Using Homebrew

* if you don't have Homebrew installed, install it following [instructions](https://brew.sh/)
* `brew install go`

Read [setting up Mac environment](a-rd600058).

## On Ubuntu

### Using Ubuntu provided pacakge

```bash
$ sudo apt-get update
$ sudo apt-get install go
```

Note that packages provided by Ubuntu are often outdated. A new version of Go is released every 6 months and Ubuntu distribution moves at slower pace.

For that reason we recommend installing binary packages.

### Using binary packages

Those instructions work on pretty much every Linux distribution:

```bash
$ sudo apt-get update
$ sudo apt-get install -y build-essential git curl wget
$ wget https://storage.googleapis.com/golang/go<versions>.gz
```

You can find the version lists [here](https://golang.org/doc/install).

```bash
# To install go1.9.3 use
$ wget https://storage.googleapis.com/golang/go1.9.3.linux-amd64.tar.gz

# Untar the file
$ sudo tar -C /usr/local -xzf go1.9.3.linux-amd64.tar.gz
$ sudo chown -R $USER:$USER /usr/local/go
$ rm go1.9.3.linux-amd64.tar.gz
```

Read [setting up Unix environment](a-rd6000f2).

## Other OSes

Follow [official instructions](https://golang.org/doc/install).
