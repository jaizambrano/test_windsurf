package service

import (
	"context"
	"errors"
	"testing"

	"fruitsapi/internal/domain"
	"fruitsapi/internal/repository"
	"fruitsapi/pkg/kvs"
)

func TestFruitService_CreateFruit(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	repo := repository.NewKVSFruitRepository(client)
	service := NewFruitService(repo)
	ctx := context.Background()

	// Test cases
	tests := []struct {
		name        string
		fruitName   string
		quantity    int
		price       float64
		owner       string
		expectError bool
	}{
		{
			name:        "ValidFruit",
			fruitName:   "manzana",
			quantity:    12,
			price:       1000,
			owner:       "test",
			expectError: false,
		},
		{
			name:        "InvalidName",
			fruitName:   "manzana123",
			quantity:    12,
			price:       1000,
			owner:       "test",
			expectError: true,
		},
		{
			name:        "InvalidQuantity",
			fruitName:   "manzana",
			quantity:    0,
			price:       1000,
			owner:       "test",
			expectError: true,
		},
		{
			name:        "InvalidPrice",
			fruitName:   "manzana",
			quantity:    12,
			price:       0,
			owner:       "test",
			expectError: true,
		},
		{
			name:        "EmptyOwner",
			fruitName:   "manzana",
			quantity:    12,
			price:       1000,
			owner:       "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Action
			fruit, err := service.CreateFruit(ctx, tt.fruitName, tt.quantity, tt.price, tt.owner)

			// Assertions
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				if fruit == nil {
					t.Errorf("Expected fruit, got nil")
					return
				}

				if fruit.Name != tt.fruitName {
					t.Errorf("Expected Name %s, got %s", tt.fruitName, fruit.Name)
				}
				if fruit.Quantity != tt.quantity {
					t.Errorf("Expected Quantity %d, got %d", tt.quantity, fruit.Quantity)
				}
				if fruit.Price != tt.price {
					t.Errorf("Expected Price %f, got %f", tt.price, fruit.Price)
				}
				if fruit.Owner != tt.owner {
					t.Errorf("Expected Owner %s, got %s", tt.owner, fruit.Owner)
				}
				if fruit.Status != "comestible" {
					t.Errorf("Expected Status %s, got %s", "comestible", fruit.Status)
				}
				if fruit.ID == "" {
					t.Errorf("Expected ID to be set, got empty string")
				}
			}
		})
	}
}

func TestFruitService_GetFruitByID(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	repo := repository.NewKVSFruitRepository(client)
	service := NewFruitService(repo)
	ctx := context.Background()

	// Create a fruit first
	fruit, err := service.CreateFruit(ctx, "manzana", 12, 1000, "test")
	if err != nil {
		t.Fatalf("Failed to create fruit: %v", err)
	}

	// Test cases
	tests := []struct {
		name        string
		id          string
		expectError bool
	}{
		{
			name:        "ExistingFruit",
			id:          fruit.ID,
			expectError: false,
		},
		{
			name:        "NonExistentFruit",
			id:          "non-existent-id",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Action
			retrievedFruit, err := service.GetFruitByID(ctx, tt.id)

			// Assertions
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}

				if retrievedFruit == nil {
					t.Errorf("Expected fruit, got nil")
					return
				}

				if retrievedFruit.ID != fruit.ID {
					t.Errorf("Expected ID %s, got %s", fruit.ID, retrievedFruit.ID)
				}
				if retrievedFruit.Name != fruit.Name {
					t.Errorf("Expected Name %s, got %s", fruit.Name, retrievedFruit.Name)
				}
			}
		})
	}
}
