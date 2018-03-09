---
Title: Configuring CSV parsing and writing
Id: 192
SOId: 20508
---
CSV is somewhat ad-hoc format. It doesn't have a specification and there are many variants.

Package [`encoding/csv`](https://golang.org/pkg/encoding/csv/) supports most common CSV formats and allows tweaking reading and writing process.

## Configuring CSV Reader

After you create `csv.Reader` with `csv.NewReader()`, you can set the following fields to change the behavior.

**`Comma`**

Most CSV files use `,` to separate records but other characters are used too. If you have a file that uses `;` as a separator you can configure a reader with `r.Comma = ';'`.

**`Comment`**

If you want to treat some CSV as comments and ignore them during reading, you can set a comment character.

For example if CSV file is:

```csv
# Comment
2013-02-08,15.07,AAL
```

you can ignore comment lines by setting `r.Comment = '#'`.

By default CSV reader doesn't detect comments and will return an error trying to parse comment line.

**`FieldsPerRecord`**

Each line in a CSV file (a record) can have a different number of fields.

If you know that e.g. CSV file you're parsing always has 5 fields in each record (line) then set `r.FieldsPerRecord = 5`. `Read()` will return an error if there's a mismatch.

If you don't know how many fields there are but know that it's always the same number, use `r.FieldsPerRecord = 0`. This is a default so you don't have to do it explicitly.

In that case `csv.Reader` will use the first line to detect number of fields and will return an error if subsequent records have a different number of fields.

If you want to allow a variable number of fields per record, set `r.FieldsPerRecord = -1`.

**`LazyQuotes`**

By default false.

If true, `csv.Reader` is more lax about parsing of quoted values i.e. a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.

**`TrimLeadingSpace`**

By default false.

If true, `csv.Reader` leading white space in a field is ignored.

**`ReuseRecord`**

By default false.

If true, the `[]string` slice returned by `Read()` might be re-used across `Read()` calls.

This is faster but you have to be more careful when using the result.

## Configuring CSV writer

After you create `csv.Writer` with `csv.NewWriter()`, you can set the following fields to change the behavior.

**`Comma`**

Field delimiter, `,` by default.

**`UseCRLF`**

False by default.

If true, uses Windows style line terminator (CRLF i.e. `\r\n`).

By default uses Unix style line terminator (LF i.e. `\n`).
