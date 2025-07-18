package handler

// CreateFruitRequest represents the request body for creating a new fruit
type CreateFruitRequest struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// ErrorResponse represents a standard error response structure
type ErrorResponse struct {
	Error string `json:"error"`
}
