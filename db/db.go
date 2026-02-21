// Package db provides a thin abstraction over database connections. This
// example currently only supports MySQL, but additional storage backends
// could be added later.
package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// NewMySQLStorage opens a connection to a MySQL database using configuration
// provided by the caller. It returns an *sql.DB which can be used for
// executing queries. The function currently logs and fatal-exits on error,
// but callers could be modified to handle errors gracefully.
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
    db, err := sql.Open("mysql", cfg.FormatDSN())

    if err != nil {
        log.Fatal(err)
    }

    return db, nil
}