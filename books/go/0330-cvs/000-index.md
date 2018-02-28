---
Title: CSV
Id: 5818
---
Package [`encoding/csv`](https://golang.org/pkg/encoding/csv/) in Go standard library provides functionality for reading and writing CSV files.

## Reading records from CSV file

Let's read stock quotes from a CSV file:

@file stocks.csv sha1:72d2bd3ce332a0d9dd82e4c35fce587516234df8 goplayground:N1IPgWv9yJx limit:5

@file read_csv.go output sha1:f7c5b98b4b69d78f0bfbfd3dc561214d4b90a0bb goplayground:hHSPtbeW5px

As per Go best practices, CSV reader operates on `io.Reader` interface, which allows it to work on files, network connections, bytes in memory etc.

`Read()` method reads one CSV line at a time and returns `[]string` slice with all fields in that line and an error.

Returning `io.EOF` as an error siginifies successfully reaching end of file.

## Reading all records from CSV file

Instead of calling `Read()` in a loop, we could read all records in one call:
```go
r := csv.NewReader(f)
records, err := r.ReadAll()
if err != nil {
    log.Fatalf("r.ReadAll() failed with '%s'\n", err)
}
// records is [][]string
fmt.Printf("Read %d records\n", len(records))
```

This time we don't have to special-case `io.EOF` as `ReadAll` does that for us.

Reading all records at once is simpler but will use more memory, especially for large CSV files.

## Writing records to CSV file

Let's now write simplified stock quotes to a CSV file:

@file write_csv.go output sha1:e4cc4a9bbbdc45e551c1b96d49da6a5545c9f0f8 goplayground:KiMuRm49LPc

Error handling here is not trivial.

We need to remember to `Flush()` at the end of writing, check if `Flush()` failed with `Error()` and also check that `Close()` didn't fail.

The need to check `Close()` errors is why we didn't use a simpler `defer f.Close()`. Correctness and robustness sometimes require more  code.

Nalues that had `,` in them were quoted because comman is used as field separator.

In production code we would also delete the CSV file in case of errors. No need to keep corrupt file around.

## Writing all records to CSV file

Just like we can read all records at once, we can write multiple records at once:

```go
w := csv.NewWriter(f)
err = w.WriteAll(records)
if err != nil {
    f.Close()
    return err
}
```
