package database

import (
	"context"
	"database/sql"
	"time"
)

type EventsModel struct {
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

func (m *EventsModel) Insert(event *Events) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2, $3, $4, $5)"

	return m.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(&event.Id)

}

func (m *EventsModel) getAll() ([]*Events, error) {

}
