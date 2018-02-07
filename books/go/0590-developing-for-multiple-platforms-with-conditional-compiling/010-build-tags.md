---
Title: Build tags
Id: 26887
Score: 0
---
```go
// +build linux

package lib

var OnlyAccessibleInLinux int // Will only be compiled in Linux
```

Negate a platform by placing `!` before it:
<!-- language: lang -->
```
// +build !windows

package lib

var NotWindows int // Will be compiled in all platforms but not Windows
```

List of platforms can be specified by separating them with spaces
<!-- language: lang -->
```
// +build linux darwin plan9

package lib

var SomeUnix int // Will be compiled in linux, darwin and plan9 but not on others
```
