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

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	_ "learn-api/docs" // Import the generated docs
	"learn-api/internal/database"
	"learn-api/internal/handlers"
	"learn-api/internal/repository"
	"learn-api/internal/services"
)

func main() {
	// Connect to the database
	if err := database.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repository, service, and handler
	entityRepo := repository.NewEntityRepository()
	entityService := services.NewEntityService(entityRepo)
	entityHandler := handlers.NewEntityHandler(entityService)

	// Create Fiber app
	app := fiber.New()

	// Add logger middleware
	app.Use(logger.New())

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Serve Swagger UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// API routes
	api := app.Group("/api/v1")

	// Entity routes
	entities := api.Group("/entities")

	entities.Get("/", entityHandler.GetAllEntitiesFiber)
	entities.Post("/", entityHandler.CreateEntityFiber)
	entities.Get("/:id", entityHandler.GetEntityByIDFiber)
	entities.Put("/:id", entityHandler.UpdateEntityFiber)
	entities.Delete("/:id", entityHandler.DeleteEntityFiber)

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
