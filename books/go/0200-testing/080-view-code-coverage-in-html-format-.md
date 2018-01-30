Title: View code coverage in HTML format 
Id: 18608
Score: 1
Body:
Run `go test` as normal, yet with the `coverprofile` flag. Then use `go tool` to view the results as HTML.

```
    go test -coverprofile=c.out
    go tool cover -html=c.out
```
|======|
