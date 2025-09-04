package services_test

import (
	"testing"

	"learn-api/internal/models"
	"learn-api/internal/repository/mocks"
	"learn-api/internal/services"
	"learn-api/pkg/errors"
)

func TestCreateEntity(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.EntityRepositoryMock{}

	// Create service with mock repository
	entityService := services.NewEntityService(mockRepo)

	// Set up the mock expectation
	mockRepo.On("Create", &models.Entity{Name: "Test Entity"}).Return(nil)

	// Create entity request
	req := &models.EntityRequest{
		Name: "Test Entity",
	}

	// Call the service method
	entity, err := entityService.CreateEntity(req)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if entity == nil {
		t.Fatal("Expected entity to be created")
	}

	if entity.Name != "Test Entity" {
		t.Errorf("Expected entity name to be 'Test Entity', got '%s'", entity.Name)
	}

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}

func TestGetEntityByID(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.EntityRepositoryMock{}

	// Create service with mock repository
	entityService := services.NewEntityService(mockRepo)

	// Create a test entity
	expectedEntity := &models.Entity{
		ID:   1,
		Name: "Test Entity",
	}

	// Set up the mock expectation
	mockRepo.On("GetByID", 1).Return(expectedEntity, nil)

	// Call the service method
	entity, err := entityService.GetEntityByID(1)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if entity == nil {
		t.Fatal("Expected entity to be found")
	}

	if entity.ID != 1 {
		t.Errorf("Expected entity ID to be 1, got %d", entity.ID)
	}

	if entity.Name != "Test Entity" {
		t.Errorf("Expected entity name to be 'Test Entity', got '%s'", entity.Name)
	}

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}

func TestGetAllEntities(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.EntityRepositoryMock{}

	// Create service with mock repository
	entityService := services.NewEntityService(mockRepo)

	// Create test entities
	expectedEntities := []*models.Entity{
		{ID: 1, Name: "Entity 1"},
		{ID: 2, Name: "Entity 2"},
		{ID: 3, Name: "Entity 3"},
	}

	// Set up the mock expectation
	mockRepo.On("GetAll").Return(expectedEntities, nil)

	// Call the service method
	entities, err := entityService.GetAllEntities()

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(entities) != 3 {
		t.Errorf("Expected 3 entities, got %d", len(entities))
	}

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}

func TestUpdateEntity(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.EntityRepositoryMock{}

	// Create service with mock repository
	entityService := services.NewEntityService(mockRepo)

	// Create a test entity for the existing entity
	existingEntity := &models.Entity{
		ID:   1,
		Name: "Original Name",
	}

	// Create updated entity
	updatedEntity := &models.Entity{
		ID:   1,
		Name: "Updated Name",
	}

	// Set up the mock expectations
	mockRepo.On("GetByID", 1).Return(existingEntity, nil)
	mockRepo.On("Update", 1, updatedEntity).Return(nil)

	// Create entity request
	req := &models.EntityRequest{
		Name: "Updated Name",
	}

	// Call the service method
	entity, err := entityService.UpdateEntity(1, req)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if entity == nil {
		t.Fatal("Expected entity to be updated")
	}

	if entity.Name != "Updated Name" {
		t.Errorf("Expected entity name to be 'Updated Name', got '%s'", entity.Name)
	}

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}

func TestUpdateEntity_NotFound(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.EntityRepositoryMock{}

	// Create service with mock repository
	entityService := services.NewEntityService(mockRepo)

	// Set up the mock expectation for non-existent entity
	mockRepo.On("GetByID", 999).Return(nil, nil)

	// Create entity request
	req := &models.EntityRequest{
		Name: "Updated Name",
	}

	// Call the service method
	entity, err := entityService.UpdateEntity(999, req)

	// Assertions
	if err != errors.ErrEntityNotFound {
		t.Fatalf("Expected ErrEntityNotFound, got %v", err)
	}

	if entity != nil {
		t.Error("Expected entity to be nil for non-existent entity")
	}

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}

func TestDeleteEntity(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.EntityRepositoryMock{}

	// Create service with mock repository
	entityService := services.NewEntityService(mockRepo)

	// Create a test entity for the existing entity
	existingEntity := &models.Entity{
		ID:   1,
		Name: "Test Entity",
	}

	// Set up the mock expectations
	mockRepo.On("GetByID", 1).Return(existingEntity, nil)
	mockRepo.On("Delete", 1).Return(nil)

	// Call the service method
	err := entityService.DeleteEntity(1)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}

func TestDeleteEntity_NotFound(t *testing.T) {
	// Create a mock repository
	mockRepo := &mocks.EntityRepositoryMock{}

	// Create service with mock repository
	entityService := services.NewEntityService(mockRepo)

	// Set up the mock expectation for non-existent entity
	mockRepo.On("GetByID", 999).Return(nil, nil)

	// Call the service method
	err := entityService.DeleteEntity(999)

	// Assertions
	if err != errors.ErrEntityNotFound {
		t.Fatalf("Expected ErrEntityNotFound, got %v", err)
	}

	// Verify mock was called
	mockRepo.AssertExpectations(t)
}
