---
Title: MySQL
Id: 196
Score: 0
SOId: 28213
---
To enable MySQL, a database driver is needed. For example [github.com/go-sql-driver/mysql](http://github.com/go-sql-driver/mysql).

```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
