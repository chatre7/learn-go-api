package validation

import (
	"learn-api/pkg/errors"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Validation implements the validation interface for ValidationError
func (v ValidationError) Validation() bool {
	return true
}

// Error implements the error interface
func (v ValidationError) Error() string {
	return v.Message
}

// Validator interface for validation rules
type Validator interface {
	Validate() []ValidationError
}

// ValidateEntityRequest validates an entity request
func ValidateEntityRequest(name string) []ValidationError {
	var errors []ValidationError

	if name == "" {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "Name is required",
		})
	}

	if len(name) > 255 {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "Name must be less than 255 characters",
		})
	}

	return errors
}

// ToAPIError converts validation errors to API errors
func ToAPIError(validationErrors []ValidationError) *errors.APIError {
	if len(validationErrors) == 0 {
		return nil
	}

	details := ""
	for i, err := range validationErrors {
		if i > 0 {
			details += "; "
		}
		details += err.Field + ": " + err.Message
	}

	return &errors.APIError{
		Code:    400,
		Message: "Validation failed",
		Details: details,
	}
}