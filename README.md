# mfile - Read data from multiple sources

mfile allows to easily read files from multiple locations easily. Provide a
schema, import a provider which implements it, and you are good to go!

---

## Example

```go

import (
    "github.com/voytechnology/mfile"
    _ "github.com/voytechnology/mfile-consul"
    _ "github.com/voytechnology/mfile-vault"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    dsn, _ := mfile.ReadFile("consul://mysql_dsn")
    pass, _ := mfile.ReadFile("vault://mysql_pass")

    c, _ := mysql.ParseDSN(string(dsn))
    c.Passwd = string(pass)
    db, _ := sql.Open("mysql", c.FormatDSN())

    // ...
}
```