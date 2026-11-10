# Go PostgreSQL CRUD API

A RESTful API built with **Golang** and **PostgreSQL** for managing a book collection.

## Features

- Create a new book
- Retrieve all books
- Retrieve a book by ID
- Update book information
- Delete a book
- Input validation
- HTTP error handling
- PostgreSQL database integration

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **Driver:** pgx/v5
- **Environment Variables:** godotenv
- **Architecture:** Handlers, Repositories, Models

## Project Structure

```text
go-postgres-api/
├── database/
│   └── connection.go
├── handlers/
│   ├── book_handler.go
│   └── response.go
├── models/
│   └── book.go
├── repositories/
│   └── book_repository.go
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── sql.sql
```

## Requirements

- Go 1.27+
- PostgreSQL
- Git

## Installation

Clone the repository:

```bash
git clone https://github.com/iskandarovi/go-postgres-api.git
cd go-postgres-api
```

Install dependencies:

```bash
go mod download
```

## Database Setup

Create a PostgreSQL database:

```sql
CREATE DATABASE library_db;
```

Connect to the database and execute the SQL commands from `sql.sql` to create the required tables.

## Environment Variables

Create a `.env` file in the project root:

```bash
cp .env.example .env
```

Configure your PostgreSQL connection:

```env
DATABASE_URL=postgres://USER:PASSWORD@localhost:5432/library_db?sslmode=disable
```

Replace `USER` and `PASSWORD` with your PostgreSQL credentials.

## Run the Application

```bash
go run .
```

The server starts at:

```text
http://localhost:8080
```

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/books` | Get all books |
| GET | `/books/{id}` | Get a book by ID |
| POST | `/books` | Create a new book |
| PUT | `/books/{id}` | Update a book |
| DELETE | `/books/{id}` | Delete a book |

## Example Request

Create a book:

**POST** `/books`

```json
{
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "price": 200000
}
```

## Run Tests

```bash
go test ./...
```

## Author

**GitHub:** [@iskandarovi](https://github.com/iskandarovi)

---

Built as a practical backend development project using Go and PostgreSQL.