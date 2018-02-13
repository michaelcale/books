---
Title: Cross Compilation
Id: 1020
---
## Introduction

The Go compiler can produce binaries for many platforms, i.e. processors and systems. Unlike with most other compilers, there is no specific requirement to cross-compiling, it is as easy to use as regular compiling.

## Syntax
 - GOOS=linux GOARCH=amd64 go build

## Remarks
Supported Operating System and Architecture target combinations [(source)](https://golang.org/doc/install/source#environment)

| $GOOS | $GOARCH |
| ----- | ------- |
| android | arm |
| darwin    | 386 |
| darwin    | amd64 |
| darwin    | arm |
| darwin    | arm64 |
| dragonfly    | amd64 |
| freebsd    | 386 |
| freebsd    | amd64 |
| freebsd    | arm |
| linux    | 386 |
| linux    | amd64 |
| linux    | arm |
| linux    | arm64 |
| linux    | ppc64 |
| linux    | ppc64le |
| linux    | mips64 |
| linux    | mips64le |
| netbsd    | 386 |
| netbsd    | amd64 |
| netbsd    | arm |
| openbsd    | 386 |
| openbsd    | amd64 |
| openbsd    | arm |
| plan9    | 386 |
| plan9    | amd64 |
| solaris    | amd64 |
| windows    | 386 |
| windows    | amd64 |
