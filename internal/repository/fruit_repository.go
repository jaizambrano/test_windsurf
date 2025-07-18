package repository

import (
	"context"
	
	"fruitsapi/internal/domain"
)

// FruitRepository defines the interface for fruit storage operations
type FruitRepository interface {
	// Save stores a fruit and returns the stored fruit with its ID
	Save(ctx context.Context, fruit *domain.Fruit) (*domain.Fruit, error)
	
	// GetByID retrieves a fruit by its ID
	GetByID(ctx context.Context, id string) (*domain.Fruit, error)
}
