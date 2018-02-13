---
Title: Simple cross compilation with go build
Id: 3299
Score: 24
---
From your project directory, run the `go build` command and specify the operating system and architecture target with the `GOOS` and `GOARCH` environment variables:

Compiling for Mac (64-bit):
```sh
GOOS=darwin GOARCH=amd64 go build
```

Compiling for Windows x86 processor:
```sh
GOOS=windows GOARCH=386 go build
```

You might also want to set the filename of the output executable manually to keep track of the architecture:

```sh
GOOS=windows GOARCH=386 go build -o appname_win_x86.exe
```

From version 1.7 and onwards you can get a list of all possible GOOS and GOARCH combinations with:

```sh
go tool dist list
```
(or for easier machine consumption `go tool dist list -json`)
