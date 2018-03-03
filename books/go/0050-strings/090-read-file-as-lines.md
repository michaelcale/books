---
Title: Read text file line by line
Id: 5eq100ld
---
Often we need to read a file line by line.

## Read file into memory and split into lines

@file read_file_as_lines.go output sha1:5aefd1e7d42fc700b0d037708467351d7ae1b6fa goplayground:g0N-gfT4K2Z

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

