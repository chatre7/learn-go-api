# Learn API
[![Test and Coverage](https://github.com/chatre7/learn-go-api/actions/workflows/test-coverage.yml/badge.svg)](https://github.com/chatre7/learn-go-api/actions/workflows/test-coverage.yml)
[![codecov](https://codecov.io/gh/chatre7/learn-go-api/branch/main/graph/badge.svg)](https://codecov.io/gh/chatre7/learn-go-api)

A RESTful API built with Go, PostgreSQL, and Docker.

## Features

- CRUD operations for entities
- PostgreSQL database integration
- Comprehensive unit testing
- Docker containerization
- Clean architecture with separation of concerns

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go          # Application entry point
├── internal/
│   ├── handlers/            # HTTP request handlers
│   ├── services/            # Business logic implementations
│   ├── models/              # Data structures
│   ├── repository/          # Data access layer
│   └── database/            # Database connection utilities
├── pkg/
│   ├── errors/              # Error handling utilities
│   └── validation/          # Validation utilities
├── tests/
│   ├── e2e/                 # End-to-end tests
│   ├── handlers/            # Tests for HTTP layer
│   ├── services/            # Tests for business logic
│   └── repository/          # Tests for data access
├── Dockerfile               # Container configuration
├── docker-compose.yml       # Multi-container setup
├── init.sql                 # Database initialization
├── go.mod                   # Go module dependencies
└── README.md                # This file
```

## API Endpoints

| Method | Endpoint             | Description          |
|--------|----------------------|----------------------|
| GET    | /api/v1/entities     | Get all entities     |
| GET    | /api/v1/entities/{id}| Get entity by ID     |
| POST   | /api/v1/entities     | Create new entity    |
| PUT    | /api/v1/entities/{id}| Update entity by ID  |
| DELETE | /api/v1/entities/{id}| Delete entity by ID  |

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose
- PostgreSQL (if running without Docker)

### Running with Docker

1. Build and start the services:
   ```bash
   docker-compose up --build
   ```

2. The API will be available at `http://localhost:8080`

### Running locally

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Set up environment variables:
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=postgres
   export DB_PASSWORD=postgres
   export DB_NAME=learnapi
   ```

3. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## Testing

Run unit tests:
```bash
go test ./tests/... -v
```

Run end-to-end tests (requires the application to be running):
```bash
go test ./tests/e2e/... -v
```

See [tests/e2e/README.md](tests/e2e/README.md) for more information about end-to-end tests.

## Database Schema

The application uses a simple entities table:

```sql
CREATE TABLE entities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Error Handling

The API returns structured error responses:

```json
{
  "error": {
    "code": 404,
    "message": "Entity not found",
    "details": "The requested entity could not be found"
  }
}
```

## Validation

Request validation is performed on entity creation and update operations. The name field is required and must be less than 255 characters.