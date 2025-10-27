package main

import (
	"database/sql"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction'up' or 'down'")
	}

	direction := os.Args[1]

	db , err := sql.Open("sqlite3","./data/db")
	if err != nil{
		log.Fatal("Couldn't Connect to the database!",err)
	}
	defer db.Close()
	
	instance, err := sqlite3.WithInstance(db,&sqlite3.Config{})
}