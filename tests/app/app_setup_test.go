package app_test

import (
    "net/http"
    "testing"

    apppkg "learn-api/internal/app"
    "learn-api/internal/models"
    "learn-api/internal/services/mocks"
)

func TestNewFiberApp_HealthAndRoutes(t *testing.T) {
    // Arrange: mock service
    mockService := &mocks.EntityServiceMock{}
    mockService.On("GetAllEntities").Return([]*models.Entity{}, nil)

    // Act: build app
    app := apppkg.NewFiberApp(mockService)

    // Assert: health endpoint
    reqHealth, _ := http.NewRequest("GET", "/health", nil)
    respHealth, err := app.Test(reqHealth)
    if err != nil {
        t.Fatalf("health request failed: %v", err)
    }
    if respHealth.StatusCode != http.StatusOK {
        t.Fatalf("expected health 200, got %d", respHealth.StatusCode)
    }

    // Assert: entities route is registered and calls service
    reqEntities, _ := http.NewRequest("GET", "/api/v1/entities/", nil)
    respEntities, err := app.Test(reqEntities)
    if err != nil {
        t.Fatalf("entities request failed: %v", err)
    }
    if respEntities.StatusCode != http.StatusOK {
        t.Fatalf("expected entities 200, got %d", respEntities.StatusCode)
    }

    mockService.AssertExpectations(t)
}
