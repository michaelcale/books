---
Title: Linux setup
Id: rd6000f2
---
After installing the compiler you need to configure [`GOPATH`](a-14406) environment variable.

```sh
mkdir -p $HOME/go
```

Add following two lines at the end of the `~/.bashrc` file

```sh
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
```

```sh
$ source ~/.bashrc
```

Test the setup by running `go version`.

You should understand the [effect of GOPATH](a-14406).
