package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-oci8" // Oracle driver
)

// CreatetOracleConn --
func CreatetOracleConn(connString string, maxIdle, maxOpen int) (oracledb *sql.DB, err error) {
	oracledb, err = sql.Open("oci8", connString)
	if err != nil {
		return nil, err
	}
	if err = oracledb.Ping(); err != nil {
		return nil, err
	}
	if maxIdle < 1 {
		maxIdle = 1
	}
	if maxOpen < 1 {
		maxOpen = 1
	}
	oracledb.SetMaxIdleConns(maxIdle)
	oracledb.SetMaxOpenConns(maxOpen)
	oracledb.SetConnMaxLifetime(10 * time.Second)
	return oracledb, nil
}
