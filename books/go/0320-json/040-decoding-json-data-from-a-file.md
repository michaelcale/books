---
Title: Decoding JSON from a file
Id: 6628
---
We can decode JSON data from a file on disk or, more broadly, any `io.Reader`, like a network connection.

Let's assume we have a file called `data.json` with the following content:

@file data.json

The following example reads the file and decodes the content:

@file decode_from_file.go output noplayground

