package rdbms

import (
	"database/sql"
	"errors"
	"reflect"
)

// RDBMS is kind of RDBMS.
type RDBMS string

const (
	// Unknown means the driver is unknown.
	Unknown RDBMS = ""

	// PostgreSQL means the driver is for PostgreSQL.
	PostgreSQL = "postgres"

	// Oracle means the driver is for Oracle.
	Oracle = "oracle"

	// MySQL means the driver is for MySQL.
	MySQL = "mysql"

	// SQLite means the driver is for SQLite.
	SQLite = "sqlite"

	// SQLServer means the driver is for Microsoft SQL Server.
	SQLServer = "mssql"

	// Firebird means the driver is for Firebird.
	Firebird = "firebird"
)

var (
	// ErrGetType shows that failed to get type of the driver.
	ErrGetType = errors.New("failed to get type of driver")

	// ErrUnknown shows that the driver is unknown.
	ErrUnknown = errors.New("unknown driver")
)

// Detect discern RDBMS from the driver.
func Detect(db *sql.DB) (RDBMS, error) {
	typ := reflect.TypeOf(db.Driver())
	if typ == nil {
		return "", ErrGetType
	}
	path := typ.PkgPath()
	return find(path)
}

func find(path string) (RDBMS, error) {
	// FIXME: replace by better search algorithm
	for _, d := range knownDrivers {
		if d.match(path) {
			return d.RDBMS, nil
		}
	}
	return "", ErrUnknown
}
