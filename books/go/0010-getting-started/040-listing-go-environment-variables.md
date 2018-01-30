Title: Listing Go Environment Variables
Id: 14405
Score: 20
Body:
Environment variables that affect the `go` tool can be viewed via the `go env [var ...]` command:

    $ go env
    GOARCH="amd64"
    GOBIN="/home/yourname/bin"
    GOEXE=""
    GOHOSTARCH="amd64"
    GOHOSTOS="linux"
    GOOS="linux"
    GOPATH="/home/yourname"
    GORACE=""
    GOROOT="/usr/lib/go"
    GOTOOLDIR="/usr/lib/go/pkg/tool/linux_amd64"
    CC="gcc"
    GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build059426571=/tmp/go-build -gno-record-gcc-switches"
    CXX="g++"
    CGO_ENABLED="1"

By default it prints the list as a shell script; however, if one or more variable names are given as arguments, it prints the value of each named variable.

    $go env GOOS GOPATH
    linux
    /home/yourname
|======|
