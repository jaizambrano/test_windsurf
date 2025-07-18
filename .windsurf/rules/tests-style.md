---
trigger: always_on
---

# Testing Style Guide

## Overview

This document outlines the testing standards and practices for the Credits MIA. Following these guidelines ensures consistent, maintainable, and effective tests across the codebase.

## Mandatory

Generate correct, clean, efficient, secure, and fully functional Go tests code.

## Test Organization

### File Structure

- Test files should be named with `_test.go` suffix
- Place test files in the same package as the code they test
- Use `package_test` for black-box testing when appropriate

```go
// Example file structure
service/
    ├── offer_service.go
    └── offer_service_test.go
repository/
    ├── offer_repository.go
    └── offer_repository_test.go
```

### Test Naming

Follow this naming convention for unit tests:

**Format**: `Test` + `Method` + `Entity` + `TestType`

Examples:

- `TestSaveAdvanceAccessGroupOK`: Test for the successful case of the Save method for the AdvanceAccessGroup entity
- `TestSaveAdvanceAccessGroupError`: Test for the error case of the Save method for the AdvanceAccessGroup entity
- `TestSaveError`: Test for the error case of the Save method (without specifying an entity)

**Important note**: The entity is optional in the test name.

### Mock Organization

- Place mocks in a dedicated `mocks` directory within the package
- Naming convention for mocks:
  - Location: `core/usecase/offer/mocks`
  - Format: `mocks.EntityName`
  - Do not include the word "mock" in the interface name: `OfferUseCase` (instead of MockOfferUseCase)

```go
repository/
    ├── mocks/
    │   └── mock_offer_repository.go  // Generated mock
    ├── offer_repository.go     // Contains OfferRepository interface
    └── offer_repository_test.go
```

## Test Structure

### Unit Tests

- Follow the Arrange-Act-Assert (AAA) pattern
- Use `context.Background()` as context within unit tests
- For object validations:
  - Use `assert.NoError` when validating the non-existence of an error type object
  - Use `assert.Nil` when validating the non-existence of a non-error type object
  - Use `assert.Error` when validating the existence of an error type object
  - Use `assert.NotNil` when validating the existence of a non-error type object
- For mocking:
  - Use `mock.Anything` when using `assert.NotNil`
  - Use specific mocks with test data when you want to validate the complete struct

### Table-Driven Tests

Use table-driven tests for testing multiple cases
## Mandatory: Testing Frameworks

### Standard Library

- Use the standard `testing` package for test structure

### Testify

- Use `github.com/stretchr/testify/assert` for assertions
- Use `github.com/stretchr/testify/require` for fatal assertions
- Use `github.com/stretchr/testify/mock` for mocking

## Testing Coverage

- Aim for minimum 90% test coverage across the codebase
- Focus on testing business logic and edge cases
- Essential paths to test:
  - Happy path
  - Error handling
  - Edge cases
  - Boundary conditions

## Best Practices

1. Keep tests simple and focused on a single unit of functionality
2. Use descriptive test names that explain what is being tested
3. Avoid testing private implementation details
4. Test error paths as thoroughly as happy paths
5. Mock external dependencies for unit tests
6. Use real dependencies for integration tests
7. Use test helpers to reduce duplication
8. Reset global state between tests
9. Run tests in parallel when possible with `t.Parallel()`
10. Write tests that are deterministic and don't depend on external state