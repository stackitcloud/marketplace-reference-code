# Variables
BINARY_NAME=marketplace-reference
PYTHON=python3
PYTHONFLAGS=-v
MAIN_PACKAGE=impl.main

# Colors for terminal output
CYAN=\033[0;36m
NC=\033[0m # No Color

.PHONY: all build clean test lint fmt run deps help

# Default target
all: deps clean 

# Clean build artifacts
clean:
	@printf "$(CYAN)Cleaning...$(NC)\n"
	rm -rf build/
	rm -rf dist/
	rm -rf *.egg-info
	find . -type f -name "*.pyc" -delete
	find . -type d -name "__pycache__" -delete
	find . -type d -name ".pytest_cache" -delete
	find . -type d -name ".coverage" -delete

# Run linter
lint:
	@printf "$(CYAN)Running linter...$(NC)\n"
	$(PYTHON) -m flake8 impl
	$(PYTHON) -m mypy impl

# Format code
fmt:
	@printf "$(CYAN)Formatting code...$(NC)\n"
	$(PYTHON) -m black impl
	$(PYTHON) -m isort impl

# Run the application
run:
	@printf "$(CYAN)Running $(BINARY_NAME)...$(NC)\n"
	$(PYTHON) -m $(MAIN_PACKAGE)

# Install dependencies
deps:
	@printf "$(CYAN)Installing dependencies...$(NC)\n"
	$(PYTHON) -m pip install --upgrade pip
	$(PYTHON) -m pip install -r requirements.txt
	$(PYTHON) -m pip install -r requirements-dev.txt

# Show help
help:
	@echo "Available targets:"
	@echo "  all    - Install dependencies, clean, and build the project"
	@echo "  clean  - Remove build artifacts"
	@echo "  lint   - Run linter and type checker"
	@echo "  fmt    - Format code with black and isort"
	@echo "  run    - Run the application"
	@echo "  deps   - Install dependencies"
	@echo "  help   - Show this help message"