package main

import (
	"log"
	"os"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// The file source driver must be imported for the `file://`
	// migration path to work. The blank identifier ensures the package's
	// init() registers the driver even though we don't reference it
	// directly in the code.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nandaiqbalh/go-backend-ecom/config"
	"github.com/nandaiqbalh/go-backend-ecom/db"
)

func main() {
	  // Build the MySQL configuration from environment variables.
    db, err := db.NewMySQLStorage(mysqlCfg.Config{
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

	
	driver, err:=  mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create MySQL driver: %v", err)
	}

   m, err := migrate.NewWithDatabaseInstance(
	"file://cmd/migrate/migrations",
	"mysql",
	driver,
   ) 

   if err != nil {
	log.Fatalf("Failed to create migrate instance: %v", err)
   }

   cmd:= os.Args[(len(os.Args)-1)]

   switch cmd {
	case "up":
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}
	log.Println("Migrations applied successfully")

   case "down":
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to rollback migrations: %v", err)
	}
	log.Println("Migrations rolled back successfully")
   default:
	log.Fatalf("Unknown command: %s. Use 'up' or 'down'.", cmd)
   }
}