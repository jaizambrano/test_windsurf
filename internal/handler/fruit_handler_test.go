package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fruitsapi/internal/domain"
	"fruitsapi/internal/repository"
	"fruitsapi/internal/service"
	"fruitsapi/pkg/kvs"
)

func TestFruitHandler_CreateFruit(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	repo := repository.NewKVSFruitRepository(client)
	service := service.NewFruitService(repo)
	handler := NewFruitHandler(service)

	tests := []struct {
		name           string
		requestBody    CreateFruitRequest
		ownerHeader    string
		expectedStatus int
	}{
		{
			name: "ValidFruit",
			requestBody: CreateFruitRequest{
				Name:     "manzana",
				Quantity: 12,
				Price:    1000,
			},
			ownerHeader:    "test",
			expectedStatus: http.StatusCreated,
		},
		{
			name: "InvalidName",
			requestBody: CreateFruitRequest{
				Name:     "manzana123",
				Quantity: 12,
				Price:    1000,
			},
			ownerHeader:    "test",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidQuantity",
			requestBody: CreateFruitRequest{
				Name:     "manzana",
				Quantity: 0,
				Price:    1000,
			},
			ownerHeader:    "test",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "InvalidPrice",
			requestBody: CreateFruitRequest{
				Name:     "manzana",
				Quantity: 12,
				Price:    0,
			},
			ownerHeader:    "test",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "MissingOwner",
			requestBody: CreateFruitRequest{
				Name:     "manzana",
				Quantity: 12,
				Price:    1000,
			},
			ownerHeader:    "",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare request
			reqBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/fruits", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			
			if tt.ownerHeader != "" {
				req.Header.Set("Owner", tt.ownerHeader)
			}

			// Prepare response recorder
			recorder := httptest.NewRecorder()

			// Execute handler
			handler.CreateFruit(recorder, req)

			// Check status code
			if recorder.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, recorder.Code)
			}

			// If success, check response body
			if tt.expectedStatus == http.StatusCreated {
				var response domain.Fruit
				err := json.NewDecoder(recorder.Body).Decode(&response)
				if err != nil {
					t.Errorf("Error decoding response body: %v", err)
				}

				if response.Name != tt.requestBody.Name {
					t.Errorf("Expected name %s, got %s", tt.requestBody.Name, response.Name)
				}
				if response.Quantity != tt.requestBody.Quantity {
					t.Errorf("Expected quantity %d, got %d", tt.requestBody.Quantity, response.Quantity)
				}
				if response.Price != tt.requestBody.Price {
					t.Errorf("Expected price %f, got %f", tt.requestBody.Price, response.Price)
				}
				if response.Owner != tt.ownerHeader {
					t.Errorf("Expected owner %s, got %s", tt.ownerHeader, response.Owner)
				}
				if response.Status != "comestible" {
					t.Errorf("Expected status %s, got %s", "comestible", response.Status)
				}
				if response.ID == "" {
					t.Error("Expected ID to be set, got empty string")
				}
			}
		})
	}
}

func TestFruitHandler_GetFruitByID(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	repo := repository.NewKVSFruitRepository(client)
	service := service.NewFruitService(repo)
	handler := NewFruitHandler(service)

	// Create a fruit first
	ctx := context.Background()
	fruit, err := service.CreateFruit(ctx, "manzana", 12, 1000, "test")
	if err != nil {
		t.Fatalf("Failed to create fruit: %v", err)
	}

	tests := []struct {
		name           string
		id             string
		expectedStatus int
	}{
		{
			name:           "ExistingFruit",
			id:             fruit.ID,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "NonExistentFruit",
			id:             "non-existent-id",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare request
			req := httptest.NewRequest(http.MethodGet, "/fruits/"+tt.id, nil)

			// Prepare response recorder
			recorder := httptest.NewRecorder()

			// Execute handler
			handler.GetFruitByID(recorder, req)

			// Check status code
			if recorder.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, recorder.Code)
			}

			// If success, check response body
			if tt.expectedStatus == http.StatusOK {
				var response domain.Fruit
				err := json.NewDecoder(recorder.Body).Decode(&response)
				if err != nil {
					t.Errorf("Error decoding response body: %v", err)
				}

				if response.ID != fruit.ID {
					t.Errorf("Expected ID %s, got %s", fruit.ID, response.ID)
				}
				if response.Name != fruit.Name {
					t.Errorf("Expected name %s, got %s", fruit.Name, response.Name)
				}
				if response.Quantity != fruit.Quantity {
					t.Errorf("Expected quantity %d, got %d", fruit.Quantity, response.Quantity)
				}
				if response.Price != fruit.Price {
					t.Errorf("Expected price %f, got %f", fruit.Price, response.Price)
				}
				if response.Owner != fruit.Owner {
					t.Errorf("Expected owner %s, got %s", fruit.Owner, response.Owner)
				}
				if response.Status != fruit.Status {
					t.Errorf("Expected status %s, got %s", fruit.Status, response.Status)
				}
			}
		})
	}
}
