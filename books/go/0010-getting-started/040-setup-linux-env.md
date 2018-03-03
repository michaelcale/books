---
Title: Linux setup
Id: rd6000f2
---
After installing the compiler you need to configure [`GOPATH`](a-14406) environment variable.

Since Go 1.8, the `GOPATH` environment variable has the default value of `$HOME/go`, so you can skip setting it.

Create the go directory with `mkdir $HOME/go`.

Add the following two lines at the end of your `~/.bashrc` file

```sh
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
```

```sh
$ source ~/.bashrc
```

Test the setup by running `go version`.

You should understand the [effect of GOPATH](a-14406).
