---
Title: Read text file line by line
Id: 5eq100ld
---

Often we need to read a file line by by lines.

## Read file into memory and split into lines

```go
func ReadFileAsLines(path string) ([]string, error) {
    d, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    s := string(d)
    lines := strings.Split(s, "\n")
    return lines, nil
}
```

## Read file as lines using a scanner

TODO: write version that uses a scanner

## Normalize newlines

TODO: write example that normalizes newlines

## A detour about newlines

There are 3 common ways to represent a newline.

**Unix**: using single character LF, which is byte 10 (0x0a), represented as "\n" in Go string literal.

**Windows**: using 2 characters: CR LF, which is bytes 13 10 (0x0d, 0x0a), represented as "\r\n" in Go string literal.

**Mac OS**: using 1 character CR (byte 13 (0x0d)), represented as "\r" in Go string literal. This is the least popular.

When splitting strings into lines you have to decide how you'll handle this.

You can assume that your code will only see e.g. Unix style line ending and only handle "\n", This won't work on files with Mac ending at all and files with Windows ending will have CR character in them.

A simple way to handle multiple newline representations is to normalize newlines and then operate on normalized version.

Finally you can write code that handles all newline endings. Inevitably, such code is a bit more complicated.
