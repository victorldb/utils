package database

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres driver
)

//CreatetPostgresConn --
func CreatetPostgresConn(connString string, maxIdle, maxOpen int) (*sql.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	if maxIdle < 1 {
		maxIdle = 1
	}
	if maxOpen < 1 {
		maxOpen = 1
	}
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)
	return db, nil
}
