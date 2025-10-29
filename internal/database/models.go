package database

import "database/sql"

type Models struct {
	Users     UserModel
	Events    EventModel
	Attendees AttendeeModel
}

func NewModels(dp *sql.DB) Models {
	return Models{}
}
