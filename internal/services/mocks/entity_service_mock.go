package mocks

import (
	"learn-api/internal/models"

	"github.com/stretchr/testify/mock"
)

// EntityServiceMock is a mock implementation of the EntityService interface
type EntityServiceMock struct {
	mock.Mock
}

// CreateEntity mocks the CreateEntity method
func (m *EntityServiceMock) CreateEntity(req *models.EntityRequest) (*models.Entity, error) {
	args := m.Called(req)
	entity, ok := args.Get(0).(*models.Entity)
	if ok {
		return entity, args.Error(1)
	}
	return nil, args.Error(1)
}

// GetEntityByID mocks the GetEntityByID method
func (m *EntityServiceMock) GetEntityByID(id int) (*models.Entity, error) {
	args := m.Called(id)
	entity, ok := args.Get(0).(*models.Entity)
	if ok {
		return entity, args.Error(1)
	}
	return nil, args.Error(1)
}

// GetAllEntities mocks the GetAllEntities method
func (m *EntityServiceMock) GetAllEntities() ([]*models.Entity, error) {
	args := m.Called()
	entities, ok := args.Get(0).([]*models.Entity)
	if ok {
		return entities, args.Error(1)
	}
	return nil, args.Error(1)
}

// UpdateEntity mocks the UpdateEntity method
func (m *EntityServiceMock) UpdateEntity(id int, req *models.EntityRequest) (*models.Entity, error) {
	args := m.Called(id, req)
	entity, ok := args.Get(0).(*models.Entity)
	if ok {
		return entity, args.Error(1)
	}
	return nil, args.Error(1)
}

// DeleteEntity mocks the DeleteEntity method
func (m *EntityServiceMock) DeleteEntity(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// AssertExpectations asserts that everything was in fact called as expected
func (m *EntityServiceMock) AssertExpectations(t mock.TestingT) bool {
	return m.Mock.AssertExpectations(t)
}

// On sets up a mock expectation
func (m *EntityServiceMock) On(methodName string, arguments ...interface{}) *mock.Call {
	return m.Mock.On(methodName, arguments...)
}