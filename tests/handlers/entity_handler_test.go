package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"learn-api/internal/handlers"
	"learn-api/internal/models"
	"learn-api/internal/services/mocks"
	"learn-api/pkg/errors"
)

func TestCreateEntity(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create a request body
	entityReq := models.EntityRequest{
		Name: "Test Entity",
	}
	jsonBody, _ := json.Marshal(entityReq)

	// Create a request
	req, err := http.NewRequest("POST", "/api/v1/entities", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Set up the mock expectation
	expectedEntity := &models.Entity{
		ID:   1,
		Name: "Test Entity",
	}
	mockService.On("CreateEntity", &entityReq).Return(expectedEntity, nil)

	// Call the handler
	entityHandler.CreateEntity(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Check the response body
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected data field in response")
	}

	if data["id"] != 1.0 {
		t.Errorf("Expected entity ID to be 1, got %v", data["id"])
	}

	if data["name"] != "Test Entity" {
		t.Errorf("Expected entity name to be 'Test Entity', got %v", data["name"])
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestGetEntityByID(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create a request
	req, err := http.NewRequest("GET", "/api/v1/entities/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Set up the mock expectation
	expectedEntity := &models.Entity{
		ID:   1,
		Name: "Test Entity",
	}
	mockService.On("GetEntityByID", 1).Return(expectedEntity, nil)

	// Call the handler with a request that has the ID in the path
	req.URL.Path = "/api/v1/entities/1"
	entityHandler.GetEntityByID(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected data field in response")
	}

	if data["id"] != 1.0 {
		t.Errorf("Expected entity ID to be 1, got %v", data["id"])
	}

	if data["name"] != "Test Entity" {
		t.Errorf("Expected entity name to be 'Test Entity', got %v", data["name"])
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestGetAllEntities(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create a request
	req, err := http.NewRequest("GET", "/api/v1/entities", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Set up the mock expectation
	expectedEntities := []*models.Entity{
		{ID: 1, Name: "Entity 1"},
		{ID: 2, Name: "Entity 2"},
	}
	mockService.On("GetAllEntities").Return(expectedEntities, nil)

	// Call the handler
	entityHandler.GetAllEntities(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	data, ok := response["data"].([]interface{})
	if !ok {
		t.Fatal("Expected data field in response")
	}

	if len(data) != 2 {
		t.Errorf("Expected 2 entities, got %d", len(data))
	}

	count, ok := response["count"].(float64)
	if !ok {
		t.Fatal("Expected count field in response")
	}

	if count != 2.0 {
		t.Errorf("Expected count to be 2, got %v", count)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestUpdateEntity(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create a request body
	entityReq := models.EntityRequest{
		Name: "Updated Entity",
	}
	jsonBody, _ := json.Marshal(entityReq)

	// Create a request
	req, err := http.NewRequest("PUT", "/api/v1/entities/1", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Set up the mock expectation
	expectedEntity := &models.Entity{
		ID:   1,
		Name: "Updated Entity",
	}
	mockService.On("UpdateEntity", 1, &entityReq).Return(expectedEntity, nil)

	// Call the handler with a request that has the ID in the path
	req.URL.Path = "/api/v1/entities/1"
	entityHandler.UpdateEntity(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var response map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected data field in response")
	}

	if data["id"] != 1.0 {
		t.Errorf("Expected entity ID to be 1, got %v", data["id"])
	}

	if data["name"] != "Updated Entity" {
		t.Errorf("Expected entity name to be 'Updated Entity', got %v", data["name"])
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestDeleteEntity(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create a request
	req, err := http.NewRequest("DELETE", "/api/v1/entities/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Set up the mock expectation
	mockService.On("DeleteEntity", 1).Return(nil)

	// Call the handler with a request that has the ID in the path
	req.URL.Path = "/api/v1/entities/1"
	entityHandler.DeleteEntity(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}

func TestGetEntityByID_NotFound(t *testing.T) {
	// Create a mock service
	mockService := &mocks.EntityServiceMock{}

	// Create handler with mock service
	entityHandler := handlers.NewEntityHandler(mockService)

	// Create a request
	req, err := http.NewRequest("GET", "/api/v1/entities/999", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Set up the mock expectation for not found
	mockService.On("GetEntityByID", 999).Return(nil, errors.ErrEntityNotFound)

	// Call the handler with a request that has the ID in the path
	req.URL.Path = "/api/v1/entities/999"
	entityHandler.GetEntityByID(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}

	// Verify mock was called
	mockService.AssertExpectations(t)
}