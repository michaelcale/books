---
Title: Easy generation of XML struct definition
Id: 334
---

Manually generating struct definitions to match structure of XML file can be tedious.

You can use [chidley](https://github.com/gnewton/chidley) tool to automatically generate struct definitions based on sample XML file.

Install the tool with `go get -u github.com/gnewton/chidley`.

Run: `chidley sample.xml`.

This will print Go struct definitions to stdout.

The tool has many options. You can get a list with `chidley -h`.
