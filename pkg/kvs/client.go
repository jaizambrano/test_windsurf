package kvs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

// Client is a simple in-memory key-value store implementation
// In a real project, this would be replaced with an actual KVS client
type Client struct {
	store map[string][]byte
	mu    sync.RWMutex
}

// NewClient creates a new instance of the KVS client
func NewClient() *Client {
	return &Client{
		store: make(map[string][]byte),
	}
}

// Set stores a value with the given key
func (c *Client) Set(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshaling value: %w", err)
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = data
	return nil
}

// Get retrieves a value by its key
func (c *Client) Get(ctx context.Context, key string, target interface{}) error {
	c.mu.RLock()
	data, ok := c.store[key]
	c.mu.RUnlock()

	if !ok {
		return errors.New("key not found")
	}

	return json.Unmarshal(data, target)
}
