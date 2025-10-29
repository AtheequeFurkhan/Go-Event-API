package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction'up' or 'down'")
	}

	direction := os.Args[1]

	// ensure data directory exists and use a .sqlite3 file
	dbPath := "./data/db.sqlite3"
	if err := os.MkdirAll("./data", 0o755); err != nil {
		log.Fatalf("unable to create data directory: %v", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Couldn't connect to the database: %v", err)
	}
	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./cmd/migrate/migrations")

	if err != nil {
		log.Fatal(err)
	}

	mig, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)

	if err != nil {
		log.Fatal(err)
	}

	switch direction {
	case "up":
		if err := mig.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := mig.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		log.Fatal("Invalid Direction")
	}

}
