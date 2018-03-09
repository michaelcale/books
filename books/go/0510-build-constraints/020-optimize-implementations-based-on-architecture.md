---
Title: Optimize implementations based on architecture
Id: 275
Score: 2
SOId: 8608
---
We can optimize a simple xor function for only architectures that support unaligned reads/writes by creating two files that define the function and prefixing them with a build constraint (for an actual example of the xor code which is out of scope here, see `crypto/cipher/xor.go` in the standard library):

```go
// +build 386 amd64 s390x

package cipher

func xorBytes(dst, a, b []byte) int { /* This function uses unaligned reads / writes to optimize the operation */ }
```

and for other architectures:

```
// +build !386,!amd64,!s390x

package cipher

func xorBytes(dst, a, b []byte) int { /* This version of the function just loops and xors */ }
```

