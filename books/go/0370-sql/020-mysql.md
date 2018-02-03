Title: MySQL
Id: 28213
Score: 0
Body:
To enable MySQL, a database driver is needed. For example [github.com/go-sql-driver/mysql][1].

    import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
    )

  [1]: http://github.com/go-sql-driver/mysql
|======|
