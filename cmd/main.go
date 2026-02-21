// The main package initializes dependencies and starts the API server.
//
// Application flow:
//  1. Load configuration from the `config` package.
//  2. Establish a connection to the MySQL database using the `db` package.
//  3. Call `initStorage` to verify the connection (and optionally perform
//     migrations or table setup).
//  4. Create and run the API server, passing the database connection for
//     handlers to use.
package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/nandaiqbalh/go-backend-ecom/cmd/api"
	"github.com/nandaiqbalh/go-backend-ecom/config"
	"github.com/nandaiqbalh/go-backend-ecom/db"
)

func main() {
    // Build the MySQL configuration from environment variables.
    db, err := db.NewMySQLStorage(mysql.Config{
        User:                 config.Envs.DBUser,
        Passwd:               config.Envs.DBPassword,
        Net:                  "tcp",
        Addr:                 config.Envs.DBAddress,
        DBName:               config.Envs.DBName,
        AllowNativePasswords: true,
        ParseTime:            true,
    })

    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Verify that the database is reachable and ready to be used.
    initStorage(db)

    // Create the API server and hand off the database connection so that
    // handlers can perform queries.
    server := api.NewAPIServer(":8080", db)
    if err := server.Run(); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

// initStorage can be used to perform any startup storage tasks. Here it
// simply pings the database to ensure the connection is valid.
func initStorage(db *sql.DB) {
    err := db.Ping()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Successfully connected to the database")
}