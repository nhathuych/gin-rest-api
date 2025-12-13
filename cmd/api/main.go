package main

import (
	"database/sql"
	"gin-rest-api/internal/database"
	"gin-rest-api/internal/env"
	"log"

	_ "gin-rest-api/docs"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

// @title Go Gin Rest API
// @version 1.0
// @description A rest API in Go using Gin framework
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	dsn := env.GetEnvString("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/gin-rest-api-development?sslmode=disable")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Cannot connect to Postgres:", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Postgres not responding:", err)
	}

	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-super-secret-key"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
