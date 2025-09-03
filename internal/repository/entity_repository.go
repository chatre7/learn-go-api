package repository

import (
	"database/sql"
	"learn-api/internal/database"
	"learn-api/internal/models"
	"learn-api/pkg/errors"
)

// EntityRepository interface defines the methods for entity operations
type EntityRepository interface {
	Create(entity *models.Entity) error
	GetByID(id int) (*models.Entity, error)
	GetAll() ([]*models.Entity, error)
	Update(id int, entity *models.Entity) error
	Delete(id int) error
}

// entityRepository implements EntityRepository interface
type entityRepository struct {
	db *sql.DB
}

// NewEntityRepository creates a new entity repository
func NewEntityRepository() EntityRepository {
	return &entityRepository{
		db: database.DB,
	}
}

// Create inserts a new entity into the database
func (r *entityRepository) Create(entity *models.Entity) error {
	query := `INSERT INTO entities (name, created_at, updated_at) VALUES ($1, NOW(), NOW()) RETURNING id`
	err := r.db.QueryRow(query, entity.Name).Scan(&entity.ID)
	if err != nil {
		return err
	}
	
	// Fetch the created entity to get timestamps
	return r.db.QueryRow(`SELECT created_at, updated_at FROM entities WHERE id = $1`, entity.ID).
		Scan(&entity.CreatedAt, &entity.UpdatedAt)
}

// GetByID retrieves an entity by its ID
func (r *entityRepository) GetByID(id int) (*models.Entity, error) {
	entity := &models.Entity{}
	query := `SELECT id, name, created_at, updated_at FROM entities WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&entity.ID, &entity.Name, &entity.CreatedAt, &entity.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.ErrDatabase
	}
	return entity, nil
}

// GetAll retrieves all entities from the database
func (r *entityRepository) GetAll() ([]*models.Entity, error) {
	query := `SELECT id, name, created_at, updated_at FROM entities ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []*models.Entity
	for rows.Next() {
		entity := &models.Entity{}
		err := rows.Scan(&entity.ID, &entity.Name, &entity.CreatedAt, &entity.UpdatedAt)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

// Update modifies an existing entity in the database
func (r *entityRepository) Update(id int, entity *models.Entity) error {
	query := `UPDATE entities SET name = $1, updated_at = NOW() WHERE id = $2`
	result, err := r.db.Exec(query, entity.Name, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	// Fetch the updated entity to get timestamps
	return r.db.QueryRow(`SELECT name, created_at, updated_at FROM entities WHERE id = $1`, id).
		Scan(&entity.Name, &entity.CreatedAt, &entity.UpdatedAt)
}

// Delete removes an entity from the database
func (r *entityRepository) Delete(id int) error {
	query := `DELETE FROM entities WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}