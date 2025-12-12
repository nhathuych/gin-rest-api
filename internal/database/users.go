package database

import "database/sql"

type UserModel struct {
	DB *sql.DB // connect object to PostgreSQL
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
