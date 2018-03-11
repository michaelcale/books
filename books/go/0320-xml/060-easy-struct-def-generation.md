---
Title: Easy generation of XML struct definition
Id: 334
---

Writing struct definitions for XML parsing can be tedious.

You can use [chidley](https://github.com/gnewton/chidley) to automatically generate struct definitions from sample XML file.

Install the tool with `go get -u github.com/gnewton/chidley`.

Run: `chidley sample.xml`.

This will print Go struct definitions to standard out.

To list all program options do `chidley -h`.
