package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"learn-api/internal/database"
	"learn-api/internal/models"
	"learn-api/internal/repository"
	"learn-api/pkg/errors"

	_ "github.com/lib/pq"
)

var testDB *sql.DB
var entityRepo repository.EntityRepository

func TestMain(m *testing.M) {
	// Set up test database connection
	if err := setupTestDB(); err != nil {
		log.Printf("Test database not available: %v", err)
		os.Setenv("SKIP_REPOSITORY_TESTS", "true")
	}

	if os.Getenv("SKIP_REPOSITORY_TESTS") != "true" {
		// Create repository instance
		entityRepo = repository.NewEntityRepository()
	}

	// Run tests
	code := m.Run()

	// Clean up
	if os.Getenv("SKIP_REPOSITORY_TESTS") != "true" {
		tearDownTestDB()
	}

	os.Exit(code)
}

func setupTestDB() error {
	var err error

	// Get test database connection details from environment variables
	host := getEnv("TEST_DB_HOST", "localhost")
	port := getEnv("TEST_DB_PORT", "5432")
	user := getEnv("TEST_DB_USER", "postgres")
	password := getEnv("TEST_DB_PASSWORD", "postgres")
	dbname := getEnv("TEST_DB_NAME", "learnapi_test")

	// Create connection string
	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

	// Open database connection
	testDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	// Check if connection is successful
	err = testDB.Ping()
	if err != nil {
		testDB.Close()
		return err
	}

	// Set the global DB for testing
	database.DB = testDB

	// Create entities table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS entities (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err = testDB.Exec(createTableQuery)
	if err != nil {
		return err
	}

	// Clear any existing data
	_, err = testDB.Exec("TRUNCATE TABLE entities RESTART IDENTITY")
	if err != nil {
		return err
	}

	log.Println("Test database setup completed!")
	return nil
}

func tearDownTestDB() {
	// Clear data
	_, err := testDB.Exec("TRUNCATE TABLE entities RESTART IDENTITY")
	if err != nil {
		log.Fatal("Error truncating entities table:", err)
	}

	// Close database connection
	testDB.Close()
	log.Println("Test database teardown completed!")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func skipIfDatabaseNotAvailable(t *testing.T) {
	if os.Getenv("SKIP_REPOSITORY_TESTS") == "true" {
		t.Skip("Skipping test because database is not available")
	}
}

func TestCreateEntity(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	entity := &models.Entity{
		Name: "Test Entity",
	}

	err := entityRepo.Create(entity)
	if err != nil {
		t.Fatalf("Error creating entity: %v", err)
	}

	if entity.ID == 0 {
		t.Error("Expected entity ID to be set")
	}

	if entity.Name != "Test Entity" {
		t.Errorf("Expected entity name to be 'Test Entity', got '%s'", entity.Name)
	}

	if entity.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}

	if entity.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestGetEntityByID(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	// First create an entity
	entity := &models.Entity{
		Name: "Test Entity",
	}
	err := entityRepo.Create(entity)
	if err != nil {
		t.Fatalf("Error creating entity: %v", err)
	}

	// Now retrieve it
	retrievedEntity, err := entityRepo.GetByID(entity.ID)
	if err != nil {
		t.Fatalf("Error retrieving entity: %v", err)
	}

	if retrievedEntity == nil {
		t.Fatal("Expected entity to be found")
	}

	if retrievedEntity.ID != entity.ID {
		t.Errorf("Expected entity ID to be %d, got %d", entity.ID, retrievedEntity.ID)
	}

	if retrievedEntity.Name != entity.Name {
		t.Errorf("Expected entity name to be '%s', got '%s'", entity.Name, retrievedEntity.Name)
	}
}

func TestGetEntityByID_NotFound(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	entity, err := entityRepo.GetByID(999999) // Non-existent ID
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if entity != nil {
		t.Error("Expected entity to be nil for non-existent ID")
	}
}

func TestGetAllEntities(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	// Create a few entities
	entities := []*models.Entity{
		{Name: "Entity 1"},
		{Name: "Entity 2"},
		{Name: "Entity 3"},
	}

	for _, entity := range entities {
		err := entityRepo.Create(entity)
		if err != nil {
			t.Fatalf("Error creating entity: %v", err)
		}
	}

	// Retrieve all entities
	allEntities, err := entityRepo.GetAll()
	if err != nil {
		t.Fatalf("Error retrieving all entities: %v", err)
	}

	if len(allEntities) < 3 {
		t.Errorf("Expected at least 3 entities, got %d", len(allEntities))
	}

	// Check if our entities are in the list
	found := 0
	for _, entity := range entities {
		for _, retrieved := range allEntities {
			if entity.Name == retrieved.Name {
				found++
				break
			}
		}
	}

	if found != 3 {
		t.Errorf("Expected to find all 3 entities, found %d", found)
	}
}

func TestUpdateEntity(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	// First create an entity
	entity := &models.Entity{
		Name: "Original Name",
	}
	err := entityRepo.Create(entity)
	if err != nil {
		t.Fatalf("Error creating entity: %v", err)
	}

	// Update the entity
	entity.Name = "Updated Name"
	err = entityRepo.Update(entity.ID, entity)
	if err != nil {
		t.Fatalf("Error updating entity: %v", err)
	}

	// Retrieve the updated entity
	updatedEntity, err := entityRepo.GetByID(entity.ID)
	if err != nil {
		t.Fatalf("Error retrieving updated entity: %v", err)
	}

	if updatedEntity.Name != "Updated Name" {
		t.Errorf("Expected updated name to be 'Updated Name', got '%s'", updatedEntity.Name)
	}

	if updatedEntity.UpdatedAt.Equal(updatedEntity.CreatedAt) {
		t.Error("Expected UpdatedAt to be different from CreatedAt after update")
	}
}

func TestUpdateEntity_NotFound(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	entity := &models.Entity{
		ID:   999999, // Non-existent ID
		Name: "Test Name",
	}

	err := entityRepo.Update(entity.ID, entity)
	if err == nil {
		t.Error("Expected error for non-existent entity")
	}

	if err != errors.ErrDatabase && err != sql.ErrNoRows {
		t.Errorf("Expected ErrDatabase or ErrNoRows, got %v", err)
	}
}

func TestDeleteEntity(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	// First create an entity
	entity := &models.Entity{
		Name: "Test Entity",
	}
	err := entityRepo.Create(entity)
	if err != nil {
		t.Fatalf("Error creating entity: %v", err)
	}

	// Delete the entity
	err = entityRepo.Delete(entity.ID)
	if err != nil {
		t.Fatalf("Error deleting entity: %v", err)
	}

	// Try to retrieve the deleted entity
	deletedEntity, err := entityRepo.GetByID(entity.ID)
	if err != nil {
		t.Fatalf("Error retrieving entity after deletion: %v", err)
	}

	if deletedEntity != nil {
		t.Error("Expected entity to be deleted (nil), but found entity")
	}
}

func TestDeleteEntity_NotFound(t *testing.T) {
	skipIfDatabaseNotAvailable(t)

	err := entityRepo.Delete(999999) // Non-existent ID
	if err == nil {
		t.Error("Expected error for non-existent entity")
	}

	if err != errors.ErrDatabase && err != sql.ErrNoRows {
		t.Errorf("Expected ErrDatabase or ErrNoRows, got %v", err)
	}
}
