package mocks

import (
	"learn-api/internal/models"

	"github.com/stretchr/testify/mock"
)

// EntityRepositoryMock is a mock implementation of the EntityRepository interface
type EntityRepositoryMock struct {
	mock.Mock
}

// Create mocks the Create method
func (m *EntityRepositoryMock) Create(entity *models.Entity) error {
	args := m.Called(entity)
	return args.Error(0)
}

// GetByID mocks the GetByID method
func (m *EntityRepositoryMock) GetByID(id int) (*models.Entity, error) {
	args := m.Called(id)
	entity, ok := args.Get(0).(*models.Entity)
	if ok {
		return entity, args.Error(1)
	}
	return nil, args.Error(1)
}

// GetAll mocks the GetAll method
func (m *EntityRepositoryMock) GetAll() ([]*models.Entity, error) {
	args := m.Called()
	entities, ok := args.Get(0).([]*models.Entity)
	if ok {
		return entities, args.Error(1)
	}
	return nil, args.Error(1)
}

// Update mocks the Update method
func (m *EntityRepositoryMock) Update(id int, entity *models.Entity) error {
	args := m.Called(id, entity)
	return args.Error(0)
}

// Delete mocks the Delete method
func (m *EntityRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// AssertExpectations asserts that everything was in fact called as expected
func (m *EntityRepositoryMock) AssertExpectations(t mock.TestingT) bool {
	return m.Mock.AssertExpectations(t)
}

// On sets up a mock expectation
func (m *EntityRepositoryMock) On(methodName string, arguments ...interface{}) *mock.Call {
	return m.Mock.On(methodName, arguments...)
}