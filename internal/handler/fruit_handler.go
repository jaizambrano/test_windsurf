package handler

import (
	"encoding/json"
	"net/http"

	"fruitsapi/internal/service"
)

// FruitHandler handles HTTP requests for fruit operations
type FruitHandler struct {
	service *service.FruitService
}

// NewFruitHandler creates a new instance of FruitHandler
func NewFruitHandler(service *service.FruitService) *FruitHandler {
	return &FruitHandler{
		service: service,
	}
}

// CreateFruit handles POST /fruits requests
func (h *FruitHandler) CreateFruit(w http.ResponseWriter, r *http.Request) {
	// Check request method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get owner from header
	owner := r.Header.Get("Owner")
	if owner == "" {
		writeJSONError(w, "Owner header is required", http.StatusBadRequest)
		return
	}

	// Parse request body
	var req CreateFruitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create fruit using service
	fruit, err := h.service.CreateFruit(r.Context(), req.Name, req.Quantity, req.Price, owner)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(fruit)
}

// GetFruitByID handles GET /fruits/{id} requests
func (h *FruitHandler) GetFruitByID(w http.ResponseWriter, r *http.Request) {
	// Check request method
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from path
	// Expected path format: /fruits/{id}
	path := r.URL.Path
	if len(path) < 8 {
		writeJSONError(w, "Invalid path", http.StatusBadRequest)
		return
	}
	id := path[8:] // Skip "/fruits/"

	// Get fruit using service
	fruit, err := h.service.GetFruitByID(r.Context(), id)
	if err != nil {
		writeJSONError(w, "Fruit not found", http.StatusNotFound)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fruit)
}

// Helper function to write JSON error responses
func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
