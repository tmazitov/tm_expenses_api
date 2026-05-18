# Expense Tracker API

A RESTful backend for tracking personal expenses, built with Go and Fiber v3.

## Features

### Authentication
- **Google OAuth2** — sign in with a Google ID token, receive a JWT access/refresh token pair
- **JWT token refresh** — exchange a valid refresh token for a new token pair
- **Logout** — revoke a refresh token and invalidate the session

### User
- **Profile** — retrieve the authenticated user's profile information

### Expenses
- **Create expense** — record a new expense with a name, price, date, and optional category
- **List expenses** — paginated list filtered by date, name, and/or category
- **Expense statistics** — aggregated spending data by weekly or monthly periods with per-category breakdowns

### Categories
- **Create category** — define a custom spending category with a color label
- **List categories** — retrieve all categories belonging to the authenticated user

### Infrastructure
- Automatic database migrations on startup via [Goose](https://github.com/pressly/goose)
- Redis-backed token storage for refresh token revocation
- Swagger UI available at `/swagger/`
- CORS support configurable via environment variables

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.25 |
| Web framework | [Fiber v3](https://github.com/gofiber/fiber) |
| Database | PostgreSQL + [Bun ORM](https://bun.uptrace.dev/) |
| Cache / token store | Redis |
| Auth | JWT ([golang-jwt/jwt v5](https://github.com/golang-jwt/jwt)) + Google OAuth2 |
| Migrations | [Goose v3](https://github.com/pressly/goose) |
| API docs | Swagger ([swaggo/swag](https://github.com/swaggo/swag)) |
| Task runner | [Task](https://taskfile.dev/) |

---

## Prerequisites

- Go 1.25+
- PostgreSQL
- Redis
- [Task](https://taskfile.dev/installation/) (`go install github.com/go-task/task/v3/cmd/task@latest`)
- [Goose](https://github.com/pressly/goose) (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/tmazitov/tm_expenses_api.git
cd tm_expenses_api
```

### 2. Install Go dependencies

```bash
go mod download
```

### 3. Configure environment variables

Create a `.env` file in the project root:

```dotenv
# PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=expense_client
DB_PASSWORD=expense_client
DB_NAME=expense_db
DB_SSL_MODE=disable

# Redis
CACHE_ADDR=localhost:6379
CACHE_DB=0

# JWT
JWT_SECRET=your_secret_key_here
JWT_ACCESS_TTL=15      # minutes
JWT_REFRESH_TTL=15     # days

# Google OAuth
GOOGLE_OAUTH_CLIENT_ID=your_google_client_id_here
```

### 4. Set up the database

Run the interactive database setup script. It will create the PostgreSQL user, database, and schema:

```bash
task init_db
```


### 5. Run the server

```bash
task dev
```

The API will be available at `http://localhost:8080`.  
Swagger UI is available at `http://localhost:8080/swagger/`.

---

## API Overview

| Method | Path | Auth | Description |
|---|---|---|---|
| `POST` | `/auth/google` | — | Authenticate with Google OAuth |
| `POST` | `/auth/refresh` | — | Refresh access token |
| `GET` | `/user` | JWT | Get user profile |
| `POST` | `/user/logout` | JWT | Logout and revoke refresh token |
| `POST` | `/expense` | JWT | Create a new expense |
| `GET` | `/expense` | JWT | List expenses |
| `GET` | `/expense/stats` | JWT | Get spending statistics |
| `POST` | `/category` | JWT | Create a category |
| `GET` | `/category` | JWT | List categories |

---

## Available Tasks

| Command | Description |
|---|---|
| `task dev` | Run the server in development mode |
| `task migrate_up` | Apply all pending migrations |
| `task migrate_down` | Roll back the last migration |
| `task migrate` | Create a new migration file |
| `task docs` | Regenerate Swagger docs from annotations |
| `task test` | Run the test suite |
