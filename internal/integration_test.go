package internal_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fruitsapi/internal/domain"
	"fruitsapi/internal/handler"
	"fruitsapi/internal/middleware"
	"fruitsapi/internal/repository"
	"fruitsapi/internal/service"
	"fruitsapi/pkg/kvs"
)

// applyMiddleware wraps a handler with multiple middleware
func applyMiddleware(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

// TestFruitsAPIIntegration tests the complete API flow for creating and retrieving fruits
func TestFruitsAPIIntegration(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	fruitRepo := repository.NewKVSFruitRepository(client)
	fruitService := service.NewFruitService(fruitRepo)
	fruitHandler := handler.NewFruitHandler(fruitService)

	// Router setup (simplified version of the router in main.go)
	router := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		
		if path == "/fruits" && r.Method == http.MethodPost {
			fruitHandler.CreateFruit(w, r)
			return
		}
		
		if len(path) > 8 && path[:8] == "/fruits/" && r.Method == http.MethodGet {
			fruitHandler.GetFruitByID(w, r)
			return
		}
		
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(handler.ErrorResponse{Error: "Not found"})
	})

	// Apply middleware
	handlerWithMiddleware := applyMiddleware(
		router,
		middleware.LoggingMiddleware,
		middleware.ContentTypeValidator,
		middleware.OwnerValidator,
	)

	// Create a test server
	server := httptest.NewServer(handlerWithMiddleware)
	defer server.Close()

	// Test case: Create a fruit and then retrieve it
	t.Run("CreateAndGetFruit", func(t *testing.T) {
		// Step 1: Create a fruit
		createReq := handler.CreateFruitRequest{
			Name:     "manzana",
			Quantity: 12,
			Price:    1000,
		}
		
		reqBody, err := json.Marshal(createReq)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}
		
		req, err := http.NewRequest(http.MethodPost, server.URL+"/fruits", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Owner", "test")
		
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("Expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
		}
		
		var createResp domain.Fruit
		if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		
		// Check that the created fruit has the expected values
		if createResp.Name != createReq.Name {
			t.Errorf("Expected name %s, got %s", createReq.Name, createResp.Name)
		}
		if createResp.Quantity != createReq.Quantity {
			t.Errorf("Expected quantity %d, got %d", createReq.Quantity, createResp.Quantity)
		}
		if createResp.Price != createReq.Price {
			t.Errorf("Expected price %f, got %f", createReq.Price, createResp.Price)
		}
		if createResp.Owner != "test" {
			t.Errorf("Expected owner %s, got %s", "test", createResp.Owner)
		}
		if createResp.Status != "comestible" {
			t.Errorf("Expected status %s, got %s", "comestible", createResp.Status)
		}
		if createResp.ID == "" {
			t.Error("Expected ID to be set, got empty string")
		}
		
		// Step 2: Retrieve the fruit using its ID
		getReq, err := http.NewRequest(http.MethodGet, server.URL+"/fruits/"+createResp.ID, nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		
		getResp, err := http.DefaultClient.Do(getReq)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer getResp.Body.Close()
		
		if getResp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, got %d", http.StatusOK, getResp.StatusCode)
		}
		
		var getResponse domain.Fruit
		if err := json.NewDecoder(getResp.Body).Decode(&getResponse); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		
		// Check that the retrieved fruit matches the created fruit
		if getResponse.ID != createResp.ID {
			t.Errorf("Expected ID %s, got %s", createResp.ID, getResponse.ID)
		}
		if getResponse.Name != createResp.Name {
			t.Errorf("Expected name %s, got %s", createResp.Name, getResponse.Name)
		}
		if getResponse.Quantity != createResp.Quantity {
			t.Errorf("Expected quantity %d, got %d", createResp.Quantity, getResponse.Quantity)
		}
		if getResponse.Price != createResp.Price {
			t.Errorf("Expected price %f, got %f", createResp.Price, getResponse.Price)
		}
		if getResponse.Owner != createResp.Owner {
			t.Errorf("Expected owner %s, got %s", createResp.Owner, getResponse.Owner)
		}
		if getResponse.Status != createResp.Status {
			t.Errorf("Expected status %s, got %s", createResp.Status, getResponse.Status)
		}
	})

	// Test case: Try to get a non-existent fruit
	t.Run("GetNonExistentFruit", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, server.URL+"/fruits/non-existent-id", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()
		
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("Expected status code %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})
}
