---
Title: Normalize newlines
Id: 801000v7
---

## A note about newlines

There are 3 common ways to represent a newline.

**Unix**: using single character LF, which is byte 10 (0x0a), represented as "\n" in Go string literal.

**Windows**: using 2 characters: CR LF, which is bytes 13 10 (0x0d, 0x0a), represented as "\r\n" in Go string literal.

**Mac OS**: using 1 character CR (byte 13 (0x0d)), represented as "\r" in Go string literal. This is the least popular.

When splitting strings into lines you have to decide how you'll handle this.

You can assume that your code will only see e.g. Unix style line endings and only handle "\n", but this won't work at all on files with Mac line endings and files with Windows line endings will have a CR character in them.

A simple way to handle multiple newline representations is to normalize the newlines and then operate on the normalized version.

Finally you can write code that handles all newline endings. Inevitably, such code is a bit more complicated.

## Normalize newlines

@file normalize_newlines.go sha1:e3dca9882e255dd24a96b8b2e7f3a688d731cdcf goplayground:Yo5PIBjvZ3A
