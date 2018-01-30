Title: Setup on Linux
Id: rd6000f2
Body:

After installing the compiler you need to configure [`GOPATH`](a-14406) environment variable.

```bash
mkdir -p $HOME/go
```

Add following two lines at the end of the `~/.bashrc` file

```bash
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:/usr/local/go/bin:$PATH
```

```
$ source ~/.bashrc
```

Test the setup by running `go version`.

You should understand the [effect of GOPATH](a-14406).
