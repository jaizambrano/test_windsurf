# Fruits API

A RESTful API for managing fruit inventory. This API allows users to create and retrieve fruits with specific properties.

## Table of Contents

- [Requirements](#requirements)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
- [Data Model](#data-model)
- [Validation Rules](#validation-rules)
- [Getting Started](#getting-started)
- [Running Tests](#running-tests)

## Requirements

- Go (latest version recommended)

## Project Structure

The project follows a clean architecture approach with separation of concerns:

```
├── cmd/
│   └── api/          # Application entry points
│       └── main.go   # Main server code
├── internal/
│   ├── domain/       # Business entities and validation rules
│   ├── handler/      # HTTP request handlers
│   ├── middleware/   # HTTP middleware components
│   ├── repository/   # Data access layer
│   └── service/      # Business logic layer
└── pkg/
    └── kvs/          # Key-Value Store client
```

## API Endpoints

### Create Fruit

Creates a new fruit in the inventory.

- **Endpoint:** `POST /fruits`
- **Headers:**
  - `Content-Type: application/json`
  - `Owner: <string>` (Required)
- **Request Body:**
  ```json
  {
    "name": "manzana",
    "quantity": 12,
    "price": 1000
  }
  ```
- **Response:** `201 Created`
  ```json
  {
    "id": "uuid",
    "name": "manzana",
    "quantity": 12,
    "price": 1000,
    "date_created": "2022-01-01T00:00-03:00",
    "date_last_updated": "2022-01-01T00:00-03:00",
    "owner": "test",
    "status": "comestible"
  }
  ```
- **Error Responses:**
  - `400 Bad Request`: Invalid input data or validation failure
  - `415 Unsupported Media Type`: Content-Type is not application/json

### Get Fruit by ID

Retrieves a specific fruit by its ID.

- **Endpoint:** `GET /fruits/{id}`
- **Response:** `200 OK`
  ```json
  {
    "id": "4b6ecad7-b6ca-4bee-9c36-0c54b7b2fc24",
    "name": "manzana",
    "quantity": 12,
    "price": 1000,
    "date_created": "2022-01-01T00:00-03:00",
    "date_last_updated": "2022-01-01T00:00-03:00",
    "owner": "test",
    "status": "comestible"
  }
  ```
- **Error Responses:**
  - `404 Not Found`: Fruit with the specified ID does not exist

## Data Model

### Fruit

| Field           | Type      | Description                           |
|-----------------|-----------|---------------------------------------|
| id              | string    | Unique identifier (UUID)              |
| name            | string    | Name of the fruit                     |
| quantity        | integer   | Amount of fruit available             |
| price           | float     | Price per unit                        |
| date_created    | timestamp | Creation timestamp                    |
| date_last_updated | timestamp | Last update timestamp                 |
| owner           | string    | Owner of the fruit record             |
| status          | string    | Status of the fruit (always "comestible") |

## Validation Rules

- **name**: Must be a string without numbers or special characters
- **quantity**: Must be a number greater than 0
- **price**: Must be a number greater than 0
- **owner**: Must not be empty (obtained from request header)

## Getting Started

### Running the API

1. Clone the repository
2. Navigate to the project directory
3. Run the API server:
   ```bash
   go run cmd/api/main.go
   ```
4. The API will be available at `http://localhost:8080`

### Example API Calls

#### Creating a Fruit
```bash
curl -X POST \
  http://localhost:8080/fruits \
  -H 'Content-Type: application/json' \
  -H 'Owner: test-owner' \
  -d '{
    "name": "manzana",
    "quantity": 12,
    "price": 1000
}'
```

#### Getting a Fruit
```bash
curl -X GET http://localhost:8080/fruits/{id}
```

## Running Tests

The project uses Test-Driven Development and includes comprehensive test coverage. To run all tests:

```bash
go test ./...
```

To run specific tests, for example, just the domain tests:

```bash
go test ./internal/domain/...
```
