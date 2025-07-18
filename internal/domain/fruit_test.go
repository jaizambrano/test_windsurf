package domain

import (
	"testing"
	"time"
)

func TestFruitValidate(t *testing.T) {
	tests := []struct {
		name        string
		fruit       Fruit
		expectError bool
	}{
		{
			name: "TestValidFruit",
			fruit: Fruit{
				Name:     "manzana",
				Quantity: 12,
				Price:    1000,
				Owner:    "test",
			},
			expectError: false,
		},
		{
			name: "TestEmptyName",
			fruit: Fruit{
				Name:     "",
				Quantity: 12,
				Price:    1000,
				Owner:    "test",
			},
			expectError: true,
		},
		{
			name: "TestNameWithNumbers",
			fruit: Fruit{
				Name:     "manzana123",
				Quantity: 12,
				Price:    1000,
				Owner:    "test",
			},
			expectError: true,
		},
		{
			name: "TestNameWithSpecialChars",
			fruit: Fruit{
				Name:     "manzana!@#",
				Quantity: 12,
				Price:    1000,
				Owner:    "test",
			},
			expectError: true,
		},
		{
			name: "TestZeroQuantity",
			fruit: Fruit{
				Name:     "manzana",
				Quantity: 0,
				Price:    1000,
				Owner:    "test",
			},
			expectError: true,
		},
		{
			name: "TestNegativeQuantity",
			fruit: Fruit{
				Name:     "manzana",
				Quantity: -5,
				Price:    1000,
				Owner:    "test",
			},
			expectError: true,
		},
		{
			name: "TestZeroPrice",
			fruit: Fruit{
				Name:     "manzana",
				Quantity: 12,
				Price:    0,
				Owner:    "test",
			},
			expectError: true,
		},
		{
			name: "TestNegativePrice",
			fruit: Fruit{
				Name:     "manzana",
				Quantity: 12,
				Price:    -100,
				Owner:    "test",
			},
			expectError: true,
		},
		{
			name: "TestEmptyOwner",
			fruit: Fruit{
				Name:     "manzana",
				Quantity: 12,
				Price:    1000,
				Owner:    "",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fruit.Validate()
			if (err != nil) != tt.expectError {
				t.Errorf("Fruit.Validate() error = %v, expectError %v", err, tt.expectError)
			}
		})
	}
}

func TestNewFruit(t *testing.T) {
	id := "test-id"
	name := "manzana"
	quantity := 12
	price := float64(1000)
	owner := "test-owner"

	fruit := NewFruit(id, name, quantity, price, owner)

	if fruit.ID != id {
		t.Errorf("NewFruit() ID = %v, want %v", fruit.ID, id)
	}
	if fruit.Name != name {
		t.Errorf("NewFruit() Name = %v, want %v", fruit.Name, name)
	}
	if fruit.Quantity != quantity {
		t.Errorf("NewFruit() Quantity = %v, want %v", fruit.Quantity, quantity)
	}
	if fruit.Price != price {
		t.Errorf("NewFruit() Price = %v, want %v", fruit.Price, price)
	}
	if fruit.Owner != owner {
		t.Errorf("NewFruit() Owner = %v, want %v", fruit.Owner, owner)
	}
	if fruit.Status != "comestible" {
		t.Errorf("NewFruit() Status = %v, want %v", fruit.Status, "comestible")
	}

	// Check that dates were set
	if fruit.DateCreated.IsZero() {
		t.Errorf("NewFruit() DateCreated should not be zero")
	}
	if fruit.DateLastUpdated.IsZero() {
		t.Errorf("NewFruit() DateLastUpdated should not be zero")
	}

	// Dates should be close to now
	now := time.Now()
	if now.Sub(fruit.DateCreated) > time.Second*5 {
		t.Errorf("NewFruit() DateCreated is too far from current time")
	}
	if now.Sub(fruit.DateLastUpdated) > time.Second*5 {
		t.Errorf("NewFruit() DateLastUpdated is too far from current time")
	}
}
