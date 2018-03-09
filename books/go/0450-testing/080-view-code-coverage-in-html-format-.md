---
Title: View code coverage in HTML format
Id: 251
Score: 1
SOId: 18608
---
Run `go test` as normal, yet with the `coverprofile` flag. Then use `go tool` to view the results as HTML.

```sh
go test -coverprofile=c.out
go tool cover -html=c.out
```

