package database

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb" // mssql driver
)

// CreatetMSsqlConn --
func CreatetMSsqlConn(connString string, maxIdle, maxOpen int) (mssqldb *sql.DB, err error) {
	mssqldb, err = sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}
	if err = mssqldb.Ping(); err != nil {
		return nil, err
	}
	if maxIdle < 1 {
		maxIdle = 1
	}
	if maxOpen < 1 {
		maxOpen = 1
	}
	mssqldb.SetMaxIdleConns(maxIdle)
	mssqldb.SetMaxOpenConns(maxOpen)
	return mssqldb, nil
}
