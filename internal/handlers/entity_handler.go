package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

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

// GetAllEntitiesFiber handles GET /api/v1/entities request for Fiber
// @Summary List all entities
// @Description Get a list of all entities
// @Tags entities
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /entities [get]
func (h *EntityHandler) GetAllEntitiesFiber(c *fiber.Ctx) error {
	entities, err := h.service.GetAllEntities()
	if err != nil {
		apiErr := errors.HandleError(err)
		return c.Status(apiErr.Code).JSON(fiber.Map{
			"error": apiErr,
		})
	}

	return c.JSON(fiber.Map{
		"data":  entities,
		"count": len(entities),
	})
}

// CreateEntityFiber handles POST /api/v1/entities request for Fiber
// @Summary Create an entity
// @Description Create a new entity with the provided data
// @Tags entities
// @Accept json
// @Produce json
// @Param entity body models.EntityRequest true "Entity to create"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /entities [post]
func (h *EntityHandler) CreateEntityFiber(c *fiber.Ctx) error {
	var req models.EntityRequest
	if err := c.BodyParser(&req); err != nil {
		err := errors.ErrInvalidRequest
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	// Validate request
	validationErrors := validation.ValidateEntityRequest(req.Name)
	if len(validationErrors) > 0 {
		err := validation.ToAPIError(validationErrors)
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	entity, err := h.service.CreateEntity(&req)
	if err != nil {
		apiErr := errors.HandleError(err)
		return c.Status(apiErr.Code).JSON(fiber.Map{
			"error": apiErr,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": entity,
	})
}

// GetEntityByIDFiber handles GET /api/v1/entities/:id request for Fiber
// @Summary Get entity by ID
// @Description Get an entity by its ID
// @Tags entities
// @Produce json
// @Param id path int true "Entity ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /entities/{id} [get]
func (h *EntityHandler) GetEntityByIDFiber(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err := errors.ErrInvalidRequest
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	entity, err := h.service.GetEntityByID(id)
	if err != nil {
		apiErr := errors.HandleError(err)
		return c.Status(apiErr.Code).JSON(fiber.Map{
			"error": apiErr,
		})
	}

	if entity == nil {
		err := errors.ErrEntityNotFound
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"data": entity,
	})
}

// UpdateEntityFiber handles PUT /api/v1/entities/:id request for Fiber
// @Summary Update entity by ID
// @Description Update an existing entity with the provided data
// @Tags entities
// @Accept json
// @Produce json
// @Param id path int true "Entity ID"
// @Param entity body models.EntityRequest true "Entity data to update"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /entities/{id} [put]
func (h *EntityHandler) UpdateEntityFiber(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err := errors.ErrInvalidRequest
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	var req models.EntityRequest
	if err := c.BodyParser(&req); err != nil {
		err := errors.ErrInvalidRequest
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	// Validate request
	validationErrors := validation.ValidateEntityRequest(req.Name)
	if len(validationErrors) > 0 {
		err := validation.ToAPIError(validationErrors)
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	entity, err := h.service.UpdateEntity(id, &req)
	if err != nil {
		apiErr := errors.HandleError(err)
		return c.Status(apiErr.Code).JSON(fiber.Map{
			"error": apiErr,
		})
	}

	if entity == nil {
		err := errors.ErrEntityNotFound
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"data": entity,
	})
}

// DeleteEntityFiber handles DELETE /api/v1/entities/:id request for Fiber
// @Summary Delete entity by ID
// @Description Delete an entity by its ID
// @Tags entities
// @Produce json
// @Param id path int true "Entity ID"
// @Success 204 "Entity deleted successfully"
// @Failure 404 {object} map[string]interface{}
// @Router /entities/{id} [delete]
func (h *EntityHandler) DeleteEntityFiber(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		err := errors.ErrInvalidRequest
		return c.Status(err.Code).JSON(fiber.Map{
			"error": err,
		})
	}

	err = h.service.DeleteEntity(id)
	if err != nil {
		apiErr := errors.HandleError(err)
		return c.Status(apiErr.Code).JSON(fiber.Map{
			"error": apiErr,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
