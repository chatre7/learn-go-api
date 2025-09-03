package errors

import (
	"database/sql"
	"errors"
	"net/http"
)

// APIError represents a structured API error
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// Common API errors
var (
	ErrEntityNotFound = &APIError{
		Code:    http.StatusNotFound,
		Message: "Entity not found",
		Details: "The requested entity could not be found",
	}

	ErrInvalidRequest = &APIError{
		Code:    http.StatusBadRequest,
		Message: "Invalid request",
		Details: "The request body is invalid or missing required fields",
	}

	ErrDatabase = &APIError{
		Code:    http.StatusInternalServerError,
		Message: "Database error",
		Details: "An error occurred while accessing the database",
	}

	ErrValidation = &APIError{
		Code:    http.StatusBadRequest,
		Message: "Validation error",
		Details: "The request data failed validation",
	}
)

// HandleError converts errors to appropriate HTTP responses
func HandleError(err error) *APIError {
	// Check if it's already an APIError
	if apiErr, ok := err.(*APIError); ok {
		return apiErr
	}

	// Handle standard errors
	if errors.Is(err, sql.ErrNoRows) {
		return ErrEntityNotFound
	}

	// Handle validation errors
	if _, ok := err.(interface{ Validation() bool }); ok {
		return ErrValidation
	}

	// Default to internal server error
	return ErrDatabase
}