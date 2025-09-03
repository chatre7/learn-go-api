package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"learn-api/internal/models"
	"learn-api/internal/services"
	"learn-api/pkg/errors"
	"learn-api/pkg/validation"
)

// EntityHandler handles HTTP requests for entities
type EntityHandler struct {
	service services.EntityService
}

// NewEntityHandler creates a new entity handler
func NewEntityHandler(service services.EntityService) *EntityHandler {
	return &EntityHandler{
		service: service,
	}
}

// CreateEntity handles POST /api/v1/entities request
func (h *EntityHandler) CreateEntity(w http.ResponseWriter, r *http.Request) {
	var req models.EntityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := errors.ErrInvalidRequest
		h.writeErrorResponse(w, err)
		return
	}

	// Validate request
	validationErrors := validation.ValidateEntityRequest(req.Name)
	if len(validationErrors) > 0 {
		err := validation.ToAPIError(validationErrors)
		h.writeErrorResponse(w, err)
		return
	}

	entity, err := h.service.CreateEntity(&req)
	if err != nil {
		apiErr := errors.HandleError(err)
		h.writeErrorResponse(w, apiErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": entity,
	})
}

// GetEntityByID handles GET /api/v1/entities/{id} request
func (h *EntityHandler) GetEntityByID(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	idStr := r.URL.Path[len("/api/v1/entities/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err := errors.ErrInvalidRequest
		h.writeErrorResponse(w, err)
		return
	}

	entity, err := h.service.GetEntityByID(id)
	if err != nil {
		apiErr := errors.HandleError(err)
		h.writeErrorResponse(w, apiErr)
		return
	}

	if entity == nil {
		err := errors.ErrEntityNotFound
		h.writeErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": entity,
	})
}

// GetAllEntities handles GET /api/v1/entities request
func (h *EntityHandler) GetAllEntities(w http.ResponseWriter, r *http.Request) {
	entities, err := h.service.GetAllEntities()
	if err != nil {
		apiErr := errors.HandleError(err)
		h.writeErrorResponse(w, apiErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data":  entities,
		"count": len(entities),
	})
}

// UpdateEntity handles PUT /api/v1/entities/{id} request
func (h *EntityHandler) UpdateEntity(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	idStr := r.URL.Path[len("/api/v1/entities/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err := errors.ErrInvalidRequest
		h.writeErrorResponse(w, err)
		return
	}

	var req models.EntityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := errors.ErrInvalidRequest
		h.writeErrorResponse(w, err)
		return
	}

	// Validate request
	validationErrors := validation.ValidateEntityRequest(req.Name)
	if len(validationErrors) > 0 {
		err := validation.ToAPIError(validationErrors)
		h.writeErrorResponse(w, err)
		return
	}

	entity, err := h.service.UpdateEntity(id, &req)
	if err != nil {
		apiErr := errors.HandleError(err)
		h.writeErrorResponse(w, apiErr)
		return
	}

	if entity == nil {
		err := errors.ErrEntityNotFound
		h.writeErrorResponse(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": entity,
	})
}

// DeleteEntity handles DELETE /api/v1/entities/{id} request
func (h *EntityHandler) DeleteEntity(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	idStr := r.URL.Path[len("/api/v1/entities/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err := errors.ErrInvalidRequest
		h.writeErrorResponse(w, err)
		return
	}

	err = h.service.DeleteEntity(id)
	if err != nil {
		apiErr := errors.HandleError(err)
		h.writeErrorResponse(w, apiErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// writeErrorResponse writes a structured error response
func (h *EntityHandler) writeErrorResponse(w http.ResponseWriter, err *errors.APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err,
	})
}