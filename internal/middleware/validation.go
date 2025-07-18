package middleware

import (
	"encoding/json"
	"net/http"
	"regexp"

	"fruitsapi/internal/handler"
)

// ContentTypeValidator ensures that the request has the correct Content-Type header
func ContentTypeValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Only check Content-Type for POST requests with a non-empty body
		if r.Method == http.MethodPost && r.ContentLength > 0 {
			contentType := r.Header.Get("Content-Type")
			if contentType != "application/json" {
				writeJSONError(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// OwnerValidator validates that the Owner header is present for POST requests
func OwnerValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			owner := r.Header.Get("Owner")
			if owner == "" {
				writeJSONError(w, "Owner header is required", http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// FruitRequestValidator validates the fruit request body
func FruitRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost && r.URL.Path == "/fruits" {
			var req handler.CreateFruitRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				writeJSONError(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
				return
			}
			
			// Reset request body for the next handler to read
			r.Body.Close()
			r.Body = http.MaxBytesReader(w, r.Body, 1048576) // 1MB limit
			
			// Validate name - must be a string without numbers or special characters
			if req.Name == "" {
				writeJSONError(w, "name cannot be empty", http.StatusBadRequest)
				return
			}
			namePattern := regexp.MustCompile(`^[a-zA-Z\s]+$`)
			if !namePattern.MatchString(req.Name) {
				writeJSONError(w, "name must contain only letters and spaces", http.StatusBadRequest)
				return
			}
			
			// Validate quantity - must be greater than 0
			if req.Quantity <= 0 {
				writeJSONError(w, "quantity must be greater than 0", http.StatusBadRequest)
				return
			}
			
			// Validate price - must be greater than 0
			if req.Price <= 0 {
				writeJSONError(w, "price must be greater than 0", http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

// Helper function to write JSON error responses
func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(handler.ErrorResponse{Error: message})
}
