package app

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/swagger"

    "learn-api/internal/handlers"
    "learn-api/internal/services"
)

// NewFiberApp builds and configures the Fiber application.
// It accepts a `services.EntityService` to allow testing with mocks.
func NewFiberApp(entityService services.EntityService) *fiber.App {
    // Initialize handler with provided service
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

    return app
}

