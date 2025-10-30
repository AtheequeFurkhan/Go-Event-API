package database

import (
	"database/sql"
	"time"
)

type EventModel struct {
	DB *sql.DB
}

type Events struct {
	Id          int       `json:"id"`
	OwnerId     string    `json:"ownerId"`
	Name        int       `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
}
