package main

import (
	"log"
	"net/http"
	"os"

	"learn-api/internal/database"
	"learn-api/internal/handlers"
	"learn-api/internal/repository"
	"learn-api/internal/services"
)

func main() {
	// Connect to the database
	if err := database.ConnectDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repository, service, and handler
	entityRepo := repository.NewEntityRepository()
	entityService := services.NewEntityService(entityRepo)
	entityHandler := handlers.NewEntityHandler(entityService)

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Set up routes
	http.HandleFunc("/api/v1/entities", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			entityHandler.GetAllEntities(w, r)
		case http.MethodPost:
			entityHandler.CreateEntity(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/entities/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			entityHandler.GetEntityByID(w, r)
		case http.MethodPut:
			entityHandler.UpdateEntity(w, r)
		case http.MethodDelete:
			entityHandler.DeleteEntity(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}