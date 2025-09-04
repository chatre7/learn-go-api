# Learn API
[![Test and Coverage](https://github.com/chatre7/learn-go-api/actions/workflows/test-coverage.yml/badge.svg)](https://github.com/chatre7/learn-go-api/actions/workflows/test-coverage.yml)
[![codecov](https://codecov.io/gh/chatre7/learn-go-api/branch/main/graph/badge.svg)](https://codecov.io/gh/chatre7/learn-go-api)

Language: English | [ภาษาไทย](README.th.md)

A RESTful API built with Go, PostgreSQL, and Docker.

## Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Web Framework](#web-framework)
- [API Endpoints](#api-endpoints)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Database Schema](#database-schema)

## Features

- CRUD operations for entities
- PostgreSQL database integration
- Comprehensive unit testing
- Docker containerization
- Clean architecture with separation of concerns
- **Web Framework**: [Fiber](https://gofiber.io/) - An Express-inspired web framework for Go

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
│   ├── database/            # Database connection utilities
│   └── app/                 # App builder (NewFiberApp)
├── pkg/
│   ├── errors/              # Error handling utilities
│   └── validation/          # Validation utilities
├── tests/
│   ├── e2e/                 # End-to-end tests
│   ├── handlers/            # Tests for HTTP layer
│   ├── services/            # Tests for business logic
│   ├── repository/          # Tests for data access
│   └── app/                 # App wiring tests (health/routes)
├── docs/                    # Swagger documentation
├── Dockerfile               # Container configuration
├── docker-compose.yml       # Multi-container setup
├── init.sql                 # Database initialization
├── go.mod                   # Go module dependencies
└── README.md                # This file
```

## Web Framework

This project uses [Fiber](https://gofiber.io/), an Express-inspired web framework for Go. Fiber is built on top of Fasthttp, the fastest HTTP engine for Go, and is designed to ease things up for fast development with zero memory allocation and performance in mind.

### Why Fiber?

- **Fast**: Built on Fasthttp, the fastest HTTP engine for Go
- **Express-like**: Familiar API for developers coming from Node.js
- **Lightweight**: Minimal overhead and small memory footprint
- **Rich Middleware**: Built-in support for common HTTP functionalities
- **Easy Testing**: Simple testing utilities for HTTP handlers

## API Endpoints

| Method | Endpoint             | Description          |
|--------|----------------------|----------------------|
| GET    | /api/v1/entities     | Get all entities     |
| GET    | /api/v1/entities/{id}| Get entity by ID     |
| POST   | /api/v1/entities     | Create new entity    |
| PUT    | /api/v1/entities/{id}| Update entity by ID  |
| DELETE | /api/v1/entities/{id}| Delete entity by ID  |
| GET    | /swagger/*           | Swagger UI           |
| GET    | /health              | Health check         |

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

## API Documentation

The API is documented using Swagger. After starting the application, you can access the Swagger UI at:
- `http://localhost:8080/swagger/index.html`

## Testing
Run all tests with full instrumentation and coverage:
```bash
go test -short -v \
  -covermode=atomic \
  -coverpkg=./internal/...,./pkg/...,./cmd/... \
  ./... \
  -coverprofile=coverage.txt
```

View coverage summary:
```bash
go tool cover -func=coverage.txt
```

Open HTML report:
```bash
go tool cover -html=coverage.txt
```

Repository tests need PostgreSQL. Quick options:
- Using Docker Compose (recommended): `docker-compose up -d db`
- Or set env vars for a local Postgres:
  - `TEST_DB_HOST=127.0.0.1`
  - `TEST_DB_PORT=5432`
  - `TEST_DB_USER=postgres`
  - `TEST_DB_PASSWORD=postgres`
  - `TEST_DB_NAME=learnapi_test`

End-to-end tests require the API to be running at `http://localhost:8080`. To skip E2E in local runs, set `SKIP_E2E_TESTS=true`.

### Code Coverage

CI publishes coverage to Codecov using `codecov.yml` to ignore docs, mocks and external test harness code. Badge is shown above.

See `tests/e2e/README.md` for more about end-to-end tests.

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
