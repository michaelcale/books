---
Title: Basic Declaration
Id: 1254
Score: 7
---
A basic struct is declared as follows:

```go
type User struct {
    FirstName, LastName string
    Email               string
    Age                 int
}
```

Each value is called a field. Fields are usually written one per line, with the field's name preceeding its type. Consecutive fields of the same type may be combined, as `FirstName` and `LastName` in the above example.
