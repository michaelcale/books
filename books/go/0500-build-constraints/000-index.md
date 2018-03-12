---
Title: Build Constraints
Id: 273
SOId: 2595
---
## Syntax
- // +build tags

## Remarks
Build tags are used for conditionally building certain files in your code. Build tags may ignore files that you don't want build unless explicitly included, or some predefined build tags may be used to have a file only be built on a particular architecture or operating system.

Build tags may appear in any kind of source file (not just Go), but they must appear near the top of the file, preceded only by blank lines and other line comments. These rules mean that in Go files a build constraint must appear before the package clause.

A series of build tags must be followed by a blank line.
