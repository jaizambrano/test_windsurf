package domain

import (
	"errors"
	"regexp"
	"time"
)

// Fruit represents a fruit item in our inventory
type Fruit struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Quantity       int       `json:"quantity"`
	Price          float64   `json:"price"`
	DateCreated    time.Time `json:"date_created"`
	DateLastUpdated time.Time `json:"date_last_updated"`
	Owner          string    `json:"owner"`
	Status         string    `json:"status"`
}

// Validate performs validation on the fruit properties according to business rules
func (f *Fruit) Validate() error {
	// Name validation: must be a string without numbers or special characters
	if f.Name == "" {
		return errors.New("name cannot be empty")
	}
	namePattern := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !namePattern.MatchString(f.Name) {
		return errors.New("name must contain only letters and spaces")
	}

	// Quantity validation: must be greater than 0
	if f.Quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}

	// Price validation: must be greater than 0
	if f.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	// Owner validation: must not be empty
	if f.Owner == "" {
		return errors.New("owner cannot be empty")
	}

	return nil
}

// NewFruit creates a new Fruit instance with provided values and default status
func NewFruit(id, name string, quantity int, price float64, owner string) *Fruit {
	now := time.Now()
	return &Fruit{
		ID:             id,
		Name:           name,
		Quantity:       quantity,
		Price:          price,
		DateCreated:    now,
		DateLastUpdated: now,
		Owner:          owner,
		Status:         "comestible", // Default status
	}
}
