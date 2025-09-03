package models

import (
	"time"
)

// Entity represents a generic entity in the system
type Entity struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EntityRequest represents the request structure for creating/updating an entity
type EntityRequest struct {
	Name string `json:"name" binding:"required"`
}