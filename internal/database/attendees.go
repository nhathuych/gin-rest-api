package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB // connect object to PostgreSQL
}

type Attendee struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}
