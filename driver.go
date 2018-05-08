package rdbms

import "strings"

// Driver represents known driver.
type Driver struct {
	PathPrefix string
	RDBMS      RDBMS
}

func (d Driver) match(path string) bool {
	if path == d.PathPrefix {
		return true
	}
	if strings.HasPrefix(path, d.PathPrefix+"/") {
		return true
	}
	return false
}

var knownDrivers = []Driver{
	{"github.com/denisenkom/go-mssqldb", SQLServer},
	{"github.com/go-sql-driver/mysql", MySQL},
	{"github.com/gwenn/gosqlite", SQLite},
	{"github.com/jackc/pgx", PostgreSQL},
	{"github.com/jbarham/gopgsqldriver", PostgreSQL},
	{"github.com/lib/pq", PostgreSQL},
	{"github.com/mattn/go-oci8", Oracle},
	{"github.com/mattn/go-sqlite3", SQLite},
	{"github.com/minus5/gofreetds", SQLServer},
	{"github.com/mxk/go-sqlite", SQLite},
	{"github.com/nakagami/firebirdsql", Firebird},
	{"github.com/rsc/sqlite", SQLite},
	{"github.com/ziutek/mymysql", MySQL},
	{"gopkg.in/goracle.v2", Oracle},
	{"gopkg.in/rana/ora.v4", Oracle},
}
