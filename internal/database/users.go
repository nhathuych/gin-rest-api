package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB // connect object to PostgreSQL
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (um *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	return um.DB.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&user.Id)
}

func (um *UserModel) getUser(query string, args ...interface{}) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := um.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (um *UserModel) Get(id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	return um.getUser(query, id)
}

func (um *UserModel) GetByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	return um.getUser(query, email)
}

func (um *UserModel) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM users"
	rows, err := um.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
