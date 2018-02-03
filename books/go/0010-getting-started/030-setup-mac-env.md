---
Title: Mac OS setup
Id: rd600058
---
After installing the compiler you need to configure [`GOPATH`](a-14406) environment variable.

Since Go 1.8, the `GOPATH` environment variable has default value of `$HOME/go` so you can skip setting it.

Create go directory with `mkdir $HOME/go`.

Add the following to `~/.bash_profile` file:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

Load the changes with `source ~/.bash_profile` or launch a new terminal for changes to take effect.

**Explanation**

File `~/.bash_profile` is executed at startup by the default `bash` shell.

By adding commands there we make them available inside every shell session.

Adding `$GOPATH/bin` to `PATH` is a matter of convenience. When you install Go programs with `go get ...`, those programs will be available from command line.

**More configuration**

I often write Go libraries, so I like to add the following shortcut to `~/.bash_profile`:

```bash
alias cdgo="cd $GOPATH/src/github.com/kjk"
```

That way a `cdgo` will cd to a directory with my Go source code is.

You need to change `github.com/kjk` for your github account.

You should understand the [effect of GOPATH](a-14406).