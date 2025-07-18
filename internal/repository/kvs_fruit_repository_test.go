package repository

import (
	"context"
	"testing"

	"fruitsapi/internal/domain"
	"fruitsapi/pkg/kvs"
)

func TestKVSFruitRepository_Save(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	repo := NewKVSFruitRepository(client)
	ctx := context.Background()

	// Test data
	fruit := domain.NewFruit("test-id", "manzana", 12, 1000, "test")

	// Action
	savedFruit, err := repo.Save(ctx, fruit)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if savedFruit.ID != fruit.ID {
		t.Errorf("Expected ID %s, got %s", fruit.ID, savedFruit.ID)
	}
}

func TestKVSFruitRepository_GetByID(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	repo := NewKVSFruitRepository(client)
	ctx := context.Background()

	// Test data
	fruit := domain.NewFruit("test-id", "manzana", 12, 1000, "test")

	// Save first
	_, err := repo.Save(ctx, fruit)
	if err != nil {
		t.Fatalf("Failed to save fruit: %v", err)
	}

	// Action
	retrievedFruit, err := repo.GetByID(ctx, fruit.ID)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if retrievedFruit.ID != fruit.ID {
		t.Errorf("Expected ID %s, got %s", fruit.ID, retrievedFruit.ID)
	}
	if retrievedFruit.Name != fruit.Name {
		t.Errorf("Expected Name %s, got %s", fruit.Name, retrievedFruit.Name)
	}
	if retrievedFruit.Quantity != fruit.Quantity {
		t.Errorf("Expected Quantity %d, got %d", fruit.Quantity, retrievedFruit.Quantity)
	}
	if retrievedFruit.Price != fruit.Price {
		t.Errorf("Expected Price %f, got %f", fruit.Price, retrievedFruit.Price)
	}
	if retrievedFruit.Owner != fruit.Owner {
		t.Errorf("Expected Owner %s, got %s", fruit.Owner, retrievedFruit.Owner)
	}
}

func TestKVSFruitRepository_GetByID_NotFound(t *testing.T) {
	// Setup
	client := kvs.NewClient()
	repo := NewKVSFruitRepository(client)
	ctx := context.Background()

	// Action
	_, err := repo.GetByID(ctx, "non-existent-id")

	// Assertions
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}
