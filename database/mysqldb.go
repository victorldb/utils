package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// CreatetMysqlConn --
func CreatetMysqlConn(connString string, maxIdle, maxOpen int) (mysqldb *sql.DB, err error) {
	mysqldb, err = sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	if err = mysqldb.Ping(); err != nil {
		return nil, err
	}
	if maxIdle < 1 {
		maxIdle = 1
	}
	if maxOpen < 1 {
		maxOpen = 1
	}
	mysqldb.SetMaxIdleConns(maxIdle)
	mysqldb.SetMaxOpenConns(maxOpen)
	mysqldb.SetConnMaxLifetime(10 * time.Second)
	return mysqldb, nil
}
