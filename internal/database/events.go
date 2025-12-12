package database

import "database/sql"

type EventModel struct {
	DB *sql.DB // connect object to PostgreSQL
}

type Event struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId" binding:"required"`
	Name        string `json:"name" binding:"required,min=3"`
	Description string `json:"description" binding:"required,min=3"`
	Date        string `json:"date" binding:"required,datetime=2006-01-02|datetime=2006-01-02T15:04:05Z07:00"`
	Location    string `json:"location" binding:"required,min=3"`
}
