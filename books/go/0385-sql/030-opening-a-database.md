---
Title: Opening a database
Id: 28214
Score: 0
---
Opening a database is database specific, here there are examples for some databases.

Sqlite 3

```go
file := "path/to/file"
db_, err := sql.Open("sqlite3", file)
if err != nil {
    panic(err)
}
```

MySql

```go
dsn := "mysql_username:CHANGEME@tcp(localhost:3306)/dbname"
db, err := sql.Open("mysql", dsn)
if err != nil {
    panic(err)
}
```
