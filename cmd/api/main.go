package main

import (
	"database/sql"
	"go-event-api/internal/database"
	"go-event-api/internal/env"
	"log"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./data/db.sqlite3")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "XXXXXXXXXXX"),
		models:    models,
	}

	if err := app.server(); err != nil {
		log.Fatal(err)
	}

}
