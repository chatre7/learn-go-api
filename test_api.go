package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EntityRequest struct {
	Name string `json:"name"`
}

type EntityResponse struct {
	Data interface{} `json:"data"`
}

func main() {
	// Wait a moment for the service to be fully ready
	time.Sleep(2 * time.Second)

	// Test creating an entity
	entityReq := EntityRequest{
		Name: "Test Fiber Entity",
	}

	jsonBody, err := json.Marshal(entityReq)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/api/v1/entities", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status Code: %d\n", resp.StatusCode)

	var response EntityResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", response)
}
