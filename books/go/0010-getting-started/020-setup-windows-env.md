Title: Windows setup
Id: rd600086
Body:

After installing the compiler you need to configure [`GOPATH`](a-14406) environment variable.

Since Go 1.8, the `GOPATH` environment variable has default value `%USERPROFILE%/go`.

will have a default value if it is unset. It defaults to $HOME/go on Unix/Linux and  on Windows.

You should understand the [effect of GOPATH](a-14406).
