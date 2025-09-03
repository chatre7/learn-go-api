# End-to-End Tests

This directory contains end-to-end tests that test the full application stack by making actual HTTP requests to the running API.

## Prerequisites

Before running the end-to-end tests, you need to have the application running. You can start it using:

```bash
docker-compose up --build
```

The tests expect the API to be available at `http://localhost:8080`.

## Running the Tests

To run the end-to-end tests, use the following command from the project root:

```bash
go test -v ./tests/e2e/...
```

If the application is not running or the database is not available, the tests will be automatically skipped rather than failing.

## Test Structure

The end-to-end tests perform the following operations:

1. **Create Entity** - Tests the POST /api/v1/entities endpoint
2. **Get Entity by ID** - Tests the GET /api/v1/entities/{id} endpoint
3. **Get All Entities** - Tests the GET /api/v1/entities endpoint
4. **Update Entity** - Tests the PUT /api/v1/entities/{id} endpoint
5. **Delete Entity** - Tests the DELETE /api/v1/entities/{id} endpoint
6. **Get Entity by ID (Not Found)** - Tests that deleted entities are no longer accessible

## Test Cleanup

The tests automatically clean up after themselves by deleting all entities that were created during the test run. This ensures that subsequent test runs start with a clean database.

## Test Environment

The tests use the actual running application, including:
- The real database (PostgreSQL)
- The actual API endpoints
- Real HTTP requests and responses

This provides confidence that all components of the application work together correctly.

## Skipping Tests

If the application is not running or the database is not available when the tests start, all end-to-end tests will be automatically skipped. This prevents the test suite from failing when the required services are not available, while still allowing the tests to run when the environment is properly set up.