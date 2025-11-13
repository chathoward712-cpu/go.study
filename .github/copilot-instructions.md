# AI Coding Agent Instructions for go.study

## Project Overview

**go.study** is a Go learning project. This document guides AI agents in contributing effectively to this codebase.

## Getting Started

### Project Structure
- This is a Go project in the early stages of development
- Follow Go project conventions from [golang.org](https://golang.org/doc/effective_go)
- Use Go modules for dependency management

### Essential Commands
- **Build**: `go build ./...`
- **Test**: `go test ./...` or `go test -v ./...` for verbose output
- **Run tests with coverage**: `go test -cover ./...`
- **Format code**: `gofmt -w .` or use `go fmt ./...`
- **Lint**: Use `golangci-lint run` if configured
- **Module management**: `go mod tidy` to clean up dependencies

## Code Conventions

### Go Idioms
- Follow the [Effective Go](https://golang.org/doc/effective_go) guide strictly
- Use short variable names in small scopes, descriptive names for package-level exports
- Prefer explicit error handling over error wrapping (unless using Go 1.13+ `%w` verb)
- Use interfaces sparingly; design around concrete types first
- Prefer composition over embedding

### File Organization
- One package per directory
- Main functionality in the package root or logical subdirectories
- `*_test.go` files should be in the same package for unit testing
- Use `cmd/` directory for executable programs if multiple binaries are needed

### Naming
- Interface names end with `-er`: `Reader`, `Writer`, `Encoder`
- Package names are singular, lowercase: `http`, `json`, `parser`
- Exported names start with uppercase: `Context`, `NewServer()`
- Private names start with lowercase: `helper()`, `privateField`

## Testing Approach

- Write tests alongside implementation using `*_test.go` files
- Use standard `testing` package; avoid external test frameworks unless specified
- Follow the pattern: `func TestFeatureName(t *testing.T) { ... }`
- Table-driven tests for multiple scenarios
- Use `t.Parallel()` for independent tests to speed up test runs

## Git Workflow

- Keep commits focused on a single logical change
- Use descriptive commit messages: "feat: add user authentication" not "fix stuff"
- Branch names: `feature/name` or `fix/issue-number`

## Before Asking for Help

- Run `go test ./...` to ensure tests pass
- Run `go fmt ./...` before committing
- Check for unused imports and variables
- Ensure code follows Go idioms

## Project-Specific Notes

### Project Goals
This is a learning-focused Go project designed to build foundational knowledge of Go language features and idioms from scratch.

### Directory Structure & Pattern
- `cmd/hello/main.go` - Entry point demonstrating all learning modules
- `internal/basics/` - Core learning examples (types, functions, interfaces, etc.)
- `pkg/utils/` - Reusable utility functions that exemplify Go patterns
- All code includes both implementation and corresponding `*_test.go` test files

### Key Learning Patterns Demonstrated

**1. Interfaces & Polymorphism** (`internal/basics/basics.go`)
```go
type Shape interface { Area() float64 }
type Rectangle struct { width, height float64 }
func (r *Rectangle) Area() float64 { return r.width * r.height }
```

**2. Closures** (`makeCounter()` in basics.go)
- Functions returning functions with captured state
- Used extensively for factories and callbacks

**3. Table-Driven Tests** (`internal/basics/basics_test.go`, `pkg/utils/helpers_test.go`)
```go
tests := []struct {
    name     string
    input    string
    expected string
}{
    {"test case", "input", "expected"},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) { ... })
}
```

### Important Conventions
- Use `// DemoXxx()` naming for demonstration functions
- Test files use `TestXxx()` for unit tests and `BenchmarkXxx()` for performance tests
- Always implement both positive and edge-case test scenarios
- Use `float64` for math operations, compare with tolerance for equality checks

