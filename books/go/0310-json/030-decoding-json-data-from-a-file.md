---
Title: Decoding JSON from a file
Id: 185
SOId: 6628
---
We can decode JSON data from a file on disk or, more broadly, any `io.Reader`, like a network connection.

Let's assume we have a file called `data.json` with the following content:

@file data.json sha1:b41ee65b22ed3032dcd719a660c5cc38878714ff goplayground:coNDIaB0MYm

The following example reads the file and decodes the content:

@file decode_from_file.go output noplayground

