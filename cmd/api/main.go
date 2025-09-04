// Package main implements a REST API for managing entities.
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /api/v1
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
    "log"
    "os"

    _ "learn-api/docs" // Import the generated docs
    "learn-api/internal/app"
    "learn-api/internal/database"
    "learn-api/internal/repository"
    "learn-api/internal/services"
)

func main() {
    // Connect to the database
    if err := database.ConnectDB(); err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Initialize repository and service
    entityRepo := repository.NewEntityRepository()
    entityService := services.NewEntityService(entityRepo)

    // Build app with dependencies
    app := app.NewFiberApp(entityService)

    // Get port from environment variable or use default
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s", port)
    log.Fatal(app.Listen(":" + port))
}
