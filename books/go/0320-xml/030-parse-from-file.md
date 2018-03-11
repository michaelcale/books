---
Title: Parse XML from file
Id: 331
---

We can decode XML data from a file on disk.

@file parse_from_file.go output sha1:d809c63b9130cc223a4e29c67b11e56e1ef8e667 goplayground:SFO3RZwmCnq

By writing a helper function `decodeFromReader`, we can easily write wrappers that will work on files, strings or network connections.
