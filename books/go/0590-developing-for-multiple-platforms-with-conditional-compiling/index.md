---
Title: Developing for Multiple Platforms with Conditional Compiling
Id: 8599
---
Platform based conditional compiling comes in two forms in Go, one is with file suffixes and the other is with build tags.

## Syntax
* After "`// +build`", a single platform or a list can follow
* Platform can be reverted by preceding it by `!` sign
* List of space separated platforms are ORed together

## Remarks

**Caveats for build tags:**
* The `// +build` constraint must be placed at the top of the file, even before package clause.
* It must be followed by one blank line to separate from package comments.

| List of valid platforms for both build tags and file suffixes |
| -------- |
| android  |
| darwin   |
| dragonfly|
| freebsd  |
| linux    |
| netbsd   |
| openbsd  |
| plan9    |
| solaris  |
| windows  |

Refer to `$GOOS` list in https://golang.org/doc/install/source#environment for the most up-to-date platform list.
