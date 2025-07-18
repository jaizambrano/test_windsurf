package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"fruitsapi/internal/handler"
	"fruitsapi/internal/middleware"
	"fruitsapi/internal/repository"
	"fruitsapi/internal/service"
	"fruitsapi/pkg/kvs"
)

// Router handles HTTP requests and routes them to the appropriate handlers
type Router struct {
	fruitHandler *handler.FruitHandler
}

// NewRouter creates a new instance of Router
func NewRouter(fruitHandler *handler.FruitHandler) *Router {
	return &Router{
		fruitHandler: fruitHandler,
	}
}

// ServeHTTP implements the http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	// Route requests based on the path
	if path == "/fruits" && req.Method == http.MethodPost {
		r.fruitHandler.CreateFruit(w, req)
		return
	}

	if strings.HasPrefix(path, "/fruits/") && req.Method == http.MethodGet {
		r.fruitHandler.GetFruitByID(w, req)
		return
	}

	// Handle 404 for unknown routes
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not found")
}

// applyMiddleware wraps a handler with multiple middleware
func applyMiddleware(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func main() {
	// Initialize KVS client
	client := kvs.NewClient()

	// Initialize repository
	fruitRepo := repository.NewKVSFruitRepository(client)

	// Initialize service
	fruitService := service.NewFruitService(fruitRepo)

	// Initialize handler
	fruitHandler := handler.NewFruitHandler(fruitService)

	// Initialize router
	router := NewRouter(fruitHandler)

	// Apply middleware
	handlerWithMiddleware := applyMiddleware(
		router,
		middleware.LoggingMiddleware,
		middleware.ContentTypeValidator,
		middleware.OwnerValidator,
	)

	// Start HTTP server
	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, handlerWithMiddleware))
}
