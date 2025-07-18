package service

import (
	"context"
	"fmt"
	"time"
	
	"github.com/google/uuid"
	
	"fruitsapi/internal/domain"
	"fruitsapi/internal/repository"
)

// FruitService handles business logic for fruit operations
type FruitService struct {
	repo repository.FruitRepository
}

// NewFruitService creates a new instance of FruitService
func NewFruitService(repo repository.FruitRepository) *FruitService {
	return &FruitService{
		repo: repo,
	}
}

// CreateFruit validates and creates a new fruit
func (s *FruitService) CreateFruit(ctx context.Context, name string, quantity int, price float64, owner string) (*domain.Fruit, error) {
	// Generate a new UUID
	id := uuid.New().String()
	
	// Create a new fruit with provided data
	fruit := domain.NewFruit(id, name, quantity, price, owner)
	
	// Validate the fruit
	if err := fruit.Validate(); err != nil {
		return nil, fmt.Errorf("invalid fruit: %w", err)
	}
	
	// Save the fruit
	return s.repo.Save(ctx, fruit)
}

// GetFruitByID retrieves a fruit by its ID
func (s *FruitService) GetFruitByID(ctx context.Context, id string) (*domain.Fruit, error) {
	return s.repo.GetByID(ctx, id)
}
