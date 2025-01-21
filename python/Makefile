# Variables
PYTHON=python3
PYTHONFLAGS=-v
MAIN_PACKAGE=marketplace_reference_code.main

# Colors for terminal output
CYAN=\033[0;36m
NC=\033[0m # No Color

.PHONY: all clean lint fmt run deps help sync-deps

# Default target
all: lint fmt run

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
	$(PYTHON) -m flake8 marketplace_reference_code 
	$(PYTHON) -m mypy marketplace_reference_code

# Format code
fmt:
	@printf "$(CYAN)Formatting code...$(NC)\n"
	$(PYTHON) -m black marketplace_reference_code
	$(PYTHON) -m isort marketplace_reference_code

# Run the application
run:
	@printf "$(CYAN)Running $(MAIN_PACKAGE)...$(NC)\n"
	$(PYTHON) -m $(MAIN_PACKAGE)

# Sync dependencies
sync-deps:
	@printf "$(CYAN)Syncing dependencies...$(NC)\n"
	$(PYTHON) -m pip install pip-tools
	$(PYTHON) -m piptools compile --output-file=requirements.txt --strip-extras pyproject.toml 
	$(PYTHON) -m piptools compile --output-file=requirements-dev.txt pyproject.toml --extra=dev --extra=test


# Update deps target to use sync-deps
deps: sync-deps
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