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

## Iterate over lines in a file

It's more efficient to only process one line at a time, as opposed to reading the whole file into memory.

We can do that using [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner):

```go
func IterLinesInFile(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    // Scan() reads next line and returns false when reached end or error
    for scanner.Scan() {
        line := scanner.Text()
        // process the line
    }
    // check if Scan() finished because of error or because it reached end of file
    return scanner.Err()
}
```

<!-- version that uses a callback -->

## A note about newlines

There are 3 common ways to represent a newline.

**Unix**: using single character LF, which is byte 10 (0x0a), represented as "\n" in Go string literal.

**Windows**: using 2 characters: CR LF, which is bytes 13 10 (0x0d, 0x0a), represented as "\r\n" in Go string literal.

**Mac OS**: using 1 character CR (byte 13 (0x0d)), represented as "\r" in Go string literal. This is the least popular.

When splitting strings into lines you have to decide how you'll handle this.

You can assume that your code will only see e.g. Unix style line ending and only handle "\n", This won't work on files with Mac ending at all and files with Windows ending will have CR character in them.

A simple way to handle multiple newline representations is to normalize newlines and then operate on normalized version.

Finally you can write code that handles all newline endings. Inevitably, such code is a bit more complicated.

## Normalize newlines

```go
func NormalizeNewlines(d []byte) []byte {
	// replace CR LF (windows) with LF (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)
	// replace CF (mac) with LF (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)
	return d
}
```
