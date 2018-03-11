---
Title: Decoding JSON from a file
Id: 185
SOId: 6628
---

We can decode JSON data from a file on disk or, more broadly, any `io.Reader`, like a network connection.

Let's assume we have a file called `data.json` with the following content:

@file data.json

The following example reads the file and decodes the content:

@file decode_from_file.go output sha1:70fdf7f5a15715cf0a99c6805b505b8ac484cd4b goplayground:GmNPErWszVh

By writing a helper function `decodeFromReader`, we can easily write wrappers that will work on files, strings or network connections.
