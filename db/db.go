package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB

func Init() {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open(os.Getenv("DB_DRIVER"), connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established")

	runMigrations(connStr)
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Println("Error closing the database:", err)
		} else {
			fmt.Println("Database connection closed")
		}
	}
}

func runMigrations(connStr string) {
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		log.Fatal("Could not create postgres driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("Could not create migrate instance:", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		if err.Error() == "dirty database" {
			log.Println("Database is in a dirty state. Forcing clean state...")
			forceErr := m.Force(1)
			if forceErr != nil {
				log.Fatal("Could not force migration version:", forceErr)
			}
			err = m.Up()
			if err != nil && err != migrate.ErrNoChange {
				log.Fatal("Could not run up migrations after forcing clean state:", err)
			}
		} else {
			log.Fatal("Could not run up migrations:", err)
		}
	}

	fmt.Println("Migrations applied successfully!")
}
