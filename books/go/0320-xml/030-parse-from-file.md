---
Title: Parse XML from file
Id: 331
---

We can decode XML data from a file on disk.

@file parse_from_file.go output sha1:583dda60d8dd584d2e9d16466422579adf06f4be goplayground:hPVUuMEb2jt

By writing a helper function `decodeFromReader`, we can easily write wrappers that will work on files, strings or network connections.
