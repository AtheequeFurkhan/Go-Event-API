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
	OwnerId     string    `json:"ownerId" binding:"required"`
	Name        int       `json:"name" binding:"required,min=3"`
	Description string    `json:"description" binding:"required,min=10"`
	Date        time.Time `json:"date" binding:"required, datetime="2000-01-01"`
	Location    string    `json:"location" binding:"required,min=3"`
}
