# gin-rest-api

## 🚀 Development Setup

### 🔄 Using Air for Live Reload

This project utilizes **Air** for automatic live reloading during development, significantly speeding up the development feedback loop.

1. **Install Air:**

```bash
go install github.com/air-verse/air@latest
```
> **Note:** Make sure your `$GOPATH/bin` directory is included in your system's `$PATH` environment variable so the `air` command is accessible globally.

2. **Run the application with Air:**
*If this is your first time, initialize the configuration file:*
```bash
air init
```

Then run:
```bash
air
```
>Your application will compile and start. Air will automatically monitor your Go source files and restart the application whenever a change is detected and saved.

## 🔐 Environment Variables

Before running the application, create a `.env` file from the example:

```bash
cp .env.example .env
```

## 🛢️ Database Setup

### 🐳 Docker Compose Setup
Use Docker Compose to quickly spin up your local database service:

```bash
docker compose up -d
```

Create a new migration file:
```bash
docker run --rm -v $(pwd)/cmd/migrate/migrations:/migrations migrate/migrate create -ext sql -dir /migrations -seq create_users_table
```

Run migrations:
```bash
docker compose run migrate
```

### 🖥️ Local PostgreSQL Setup (Using Go Binary)

Create a new migration file:
```bash
migrate create -ext sql -dir cmd/migrate/migrations -seq create_users_table
```

Apply Migrations (Up):
```bash
go run ./cmd/migrate/main.go up
```

Rollback Migrations (Down):
```bash
go run ./cmd/migrate/main.go down
```

## 📘 Swagger API Documentation

🛠️ Generate Swagger Docs

```bash
swag init --dir ./cmd/api --output ./docs --parseDependency --parseInternal --parseDepth 1
```

🌐 Access Swagger UI
```bash
http://localhost:8080/swagger
```
