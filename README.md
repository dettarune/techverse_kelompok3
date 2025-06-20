# Todo

This is a simple RESTful API built with Go for user authentication and todo management. It features JWT-based authentication, user registration/login, and CRUD operations for todos, all backed by a PostgreSQL database.

## Features

- User registration and login with JWT authentication
- CRUD operations for todos (create, read, update, delete)
- Per-user todo isolation
- Built with Go, GORM, and PostgreSQL

## Requirements

- Go 1.24+
- PostgreSQL

## Setup

1. **Clone the repository:**
   ```sh
   git clone <repo-url>
   cd techverse
   ```
2. **Set up environment variables:**
   Copy the example file and edit it as needed:
   ```sh
   cp .env.example .env
   # Then edit .env with your own values
   ```
3. **Install dependencies:**
   ```sh
   go mod tidy
   ```
4. **Run database migrations:**
   The app will auto-migrate tables on startup.

## Build & Run

- **Build:**
  ```sh
  make build
  ```
- **Run:**
  ```sh
  make run
  ```
- **Live reload (requires [air](https://github.com/cosmtrek/air))**
  ```sh
  make run/live
  ```

The server will start on `http://localhost:8080`.

## API Documentation

API endpoints and example requests/responses are available in the provided Postman collection.
