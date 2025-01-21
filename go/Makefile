# Variables
BINARY_NAME=marketplace-reference
GO=go
GOFLAGS=-v
MAIN_PACKAGE=./cmd

# Colors for terminal output
CYAN=\033[0;36m
NC=\033[0m # No Color

.PHONY: all build clean test lint fmt vet run help

# Default target
all: clean build

# Build the application
build:
	@printf "$(CYAN)Building $(BINARY_NAME)...$(NC)\n"
	$(GO) build $(GOFLAGS) -o bin/$(BINARY_NAME) $(MAIN_PACKAGE)

# Clean build artifacts
clean:
	@printf "$(CYAN)Cleaning...$(NC)\n"
	rm -rf bin/
	$(GO) clean

# Run tests
test:
	@printf "$(CYAN)Running tests...$(NC)\n"
	$(GO) test ./internal/... ./utils/... -v

# Run linter
lint:
	@printf "$(CYAN)Running linter...$(NC)\n"
	golangci-lint run ./...

# Format code
fmt:
	@printf "$(CYAN)Formatting code...$(NC)\n"
	$(GO) fmt ./...

# Run go vet
vet:
	@printf "$(CYAN)Running go vet...$(NC)\n"
	$(GO) vet ./cmd/... ./internal/... ./utils/...

# Run the application
run:
	@printf "$(CYAN)Running $(BINARY_NAME)...$(NC)\n"
	$(GO) run $(MAIN_PACKAGE)

# Download dependencies
deps:
	@printf "$(CYAN)Downloading dependencies...$(NC)\n"
	$(GO) mod download
	$(GO) mod tidy

# Show help
help:
	@echo "Available targets:"
	@echo "  all    - Clean and build the project"
	@echo "  build  - Build the application"
	@echo "  clean  - Remove build artifacts"
	@echo "  test   - Run tests"
	@echo "  lint   - Run linter"
	@echo "  fmt    - Format code"
	@echo "  vet    - Run go vet"
	@echo "  run    - Run the application"
	@echo "  deps   - Download dependencies"
	@echo "  help   - Show this help message"