package services

import (
	"learn-api/internal/models"
	"learn-api/internal/repository"
	"learn-api/pkg/errors"
)

// EntityService interface defines the methods for entity service operations
type EntityService interface {
	CreateEntity(req *models.EntityRequest) (*models.Entity, error)
	GetEntityByID(id int) (*models.Entity, error)
	GetAllEntities() ([]*models.Entity, error)
	UpdateEntity(id int, req *models.EntityRequest) (*models.Entity, error)
	DeleteEntity(id int) error
}

// entityService implements EntityService interface
type entityService struct {
	repo repository.EntityRepository
}

// NewEntityService creates a new entity service
func NewEntityService(repo repository.EntityRepository) EntityService {
	return &entityService{
		repo: repo,
	}
}

// CreateEntity creates a new entity
func (s *entityService) CreateEntity(req *models.EntityRequest) (*models.Entity, error) {
	entity := &models.Entity{
		Name: req.Name,
	}

	err := s.repo.Create(entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// GetEntityByID retrieves an entity by its ID
func (s *entityService) GetEntityByID(id int) (*models.Entity, error) {
	return s.repo.GetByID(id)
}

// GetAllEntities retrieves all entities
func (s *entityService) GetAllEntities() ([]*models.Entity, error) {
	return s.repo.GetAll()
}

// UpdateEntity updates an existing entity
func (s *entityService) UpdateEntity(id int, req *models.EntityRequest) (*models.Entity, error) {
	// First, check if entity exists
	entity, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	if entity == nil {
		return nil, errors.ErrEntityNotFound
	}

	// Update entity fields
	entity.Name = req.Name

	err = s.repo.Update(id, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// DeleteEntity deletes an entity by its ID
func (s *entityService) DeleteEntity(id int) error {
	// First, check if entity exists
	entity, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	
	if entity == nil {
		return errors.ErrEntityNotFound
	}

	return s.repo.Delete(id)
}