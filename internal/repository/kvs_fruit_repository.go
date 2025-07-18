package repository

import (
	"context"
	"fmt"

	"fruitsapi/internal/domain"
	"fruitsapi/pkg/kvs"
)

// KVSFruitRepository implements FruitRepository using a KVS client
type KVSFruitRepository struct {
	client *kvs.Client
}

// NewKVSFruitRepository creates a new instance of KVSFruitRepository
func NewKVSFruitRepository(client *kvs.Client) *KVSFruitRepository {
	return &KVSFruitRepository{
		client: client,
	}
}

// Save stores a fruit in the KVS
func (r *KVSFruitRepository) Save(ctx context.Context, fruit *domain.Fruit) (*domain.Fruit, error) {
	if err := r.client.Set(ctx, fruit.ID, fruit); err != nil {
		return nil, fmt.Errorf("error saving fruit to KVS: %w", err)
	}
	return fruit, nil
}

// GetByID retrieves a fruit from the KVS by its ID
func (r *KVSFruitRepository) GetByID(ctx context.Context, id string) (*domain.Fruit, error) {
	var fruit domain.Fruit
	if err := r.client.Get(ctx, id, &fruit); err != nil {
		return nil, fmt.Errorf("error retrieving fruit from KVS: %w", err)
	}
	return &fruit, nil
}
