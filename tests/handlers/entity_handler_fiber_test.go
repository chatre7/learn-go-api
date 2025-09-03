package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"

	"learn-api/internal/handlers"
	"learn-api/internal/models"
	"learn-api/internal/services/mocks"
)

func TestCreateEntityFiber(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create Fiber app for testing
	app := fiber.New()
	app.Post("/entities", entityHandler.CreateEntityFiber)

	// Set up the mock expectation
	entityReq := &models.EntityRequest{
		Name: "Test Entity",
	}
	expectedEntity := &models.Entity{
		ID:   1,
		Name: "Test Entity",
	}
	mockService.On("CreateEntity", entityReq).Return(expectedEntity, nil)

	// Create request body
	body, _ := json.Marshal(entityReq)

	// Make request
	req, _ := http.NewRequest("POST", "/entities", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// Check status code
	if resp.StatusCode != fiber.StatusCreated {
		t.Errorf("Expected status code %d, got %d", fiber.StatusCreated, resp.StatusCode)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestGetEntityByIDFiber(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create Fiber app for testing
	app := fiber.New()
	app.Get("/entities/:id", entityHandler.GetEntityByIDFiber)

	// Set up the mock expectation
	expectedEntity := &models.Entity{
		ID:   1,
		Name: "Test Entity",
	}
	mockService.On("GetEntityByID", 1).Return(expectedEntity, nil)

	// Make request
	req, _ := http.NewRequest("GET", "/entities/1", nil)

	// Perform request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// Check status code
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status code %d, got %d", fiber.StatusOK, resp.StatusCode)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestGetAllEntitiesFiber(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create Fiber app for testing
	app := fiber.New()
	app.Get("/entities", entityHandler.GetAllEntitiesFiber)

	// Set up the mock expectation
	expectedEntities := []*models.Entity{
		{ID: 1, Name: "Entity 1"},
		{ID: 2, Name: "Entity 2"},
	}
	mockService.On("GetAllEntities").Return(expectedEntities, nil)

	// Make request
	req, _ := http.NewRequest("GET", "/entities", nil)

	// Perform request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// Check status code
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status code %d, got %d", fiber.StatusOK, resp.StatusCode)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestUpdateEntityFiber(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create Fiber app for testing
	app := fiber.New()
	app.Put("/entities/:id", entityHandler.UpdateEntityFiber)

	// Set up the mock expectation
	entityReq := &models.EntityRequest{
		Name: "Updated Name",
	}
	expectedEntity := &models.Entity{
		ID:   1,
		Name: "Updated Name",
	}
	mockService.On("UpdateEntity", 1, entityReq).Return(expectedEntity, nil)

	// Create request body
	body, _ := json.Marshal(entityReq)

	// Make request
	req, _ := http.NewRequest("PUT", "/entities/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// Check status code
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status code %d, got %d", fiber.StatusOK, resp.StatusCode)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestDeleteEntityFiber(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create Fiber app for testing
	app := fiber.New()
	app.Delete("/entities/:id", entityHandler.DeleteEntityFiber)

	// Set up the mock expectation
	mockService.On("DeleteEntity", 1).Return(nil)

	// Make request
	req, _ := http.NewRequest("DELETE", "/entities/1", nil)

	// Perform request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// Check status code
	if resp.StatusCode != fiber.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", fiber.StatusNoContent, resp.StatusCode)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestGetEntityByIDNotFoundFiber(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create Fiber app for testing
	app := fiber.New()
	app.Get("/entities/:id", entityHandler.GetEntityByIDFiber)

	// Set up the mock expectation for non-existent entity
	mockService.On("GetEntityByID", 999).Return(nil, nil)

	// Make request
	req, _ := http.NewRequest("GET", "/entities/999", nil)

	// Perform request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}

	// Check status code
	if resp.StatusCode != fiber.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", fiber.StatusNotFound, resp.StatusCode)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}
