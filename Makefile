# Jobot Makefile

# Variables
APP_NAME=jobot
DOCKER_COMPOSE_FILE=deploy/debug/docker-compose.yaml
DOCKER_IMAGE_NAME=jobot-app

# Go commands
.PHONY: build run test clean deps fmt lint

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	go build -o bin/$(APP_NAME) ./cmd/app

# Run the application locally
run:
	@echo "Running $(APP_NAME)..."
	go run ./cmd/app

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	go clean

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	golangci-lint run

# Docker commands
.PHONY: docker-build docker-run docker-stop docker-logs docker-clean

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run with Docker Compose
docker-up:
	@echo "Starting services with Docker Compose..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop Docker Compose services
docker-down:
	@echo "Stopping Docker Compose services..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down

# Stop and remove volumes
docker-down-volumes:
	@echo "Stopping Docker Compose services and removing volumes..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down -v

# View logs
docker-logs:
	@echo "Viewing logs..."
	docker compose -f $(DOCKER_COMPOSE_FILE) logs -f

# View logs for specific service
docker-logs-app:
	@echo "Viewing app logs..."
	docker compose -f $(DOCKER_COMPOSE_FILE) logs -f jobot-app

docker-logs-db:
	@echo "Viewing database logs..."
	docker compose -f $(DOCKER_COMPOSE_FILE) logs -f postgres

# Rebuild and restart services
docker-rebuild:
	@echo "Rebuilding and restarting services..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d --build

# Clean Docker resources
docker-clean:
	@echo "Cleaning Docker resources..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down -v --rmi all
	docker system prune -f

# API Documentation commands
.PHONY: api-docs api-docs-open api-validate

# Open API documentation
api-docs-open:
	@echo "Opening Swagger UI..."
	@echo "URL: http://localhost:8080/api/docs"
	@command -v xdg-open >/dev/null 2>&1 && xdg-open http://localhost:8080/api/docs || \
	command -v open >/dev/null 2>&1 && open http://localhost:8080/api/docs || \
	echo "Please open http://localhost:8080/api/docs in your browser"

# Show API documentation info
api-docs:
	@echo "API Documentation:"
	@echo "  Swagger UI:  http://localhost:8080/api/docs"
	@echo "  OpenAPI YAML: http://localhost:8080/api/swagger.yaml"
	@echo "  OpenAPI JSON: http://localhost:8080/api/swagger.json"
	@echo ""
	@echo "To open Swagger UI, run: make api-docs-open"

# Validate OpenAPI specification
api-validate:
	@echo "Validating OpenAPI specification..."
	@command -v swagger >/dev/null 2>&1 && swagger validate api/swagger.yaml || \
	echo "swagger CLI not installed. Install with: npm install -g swagger-cli"

# Development commands
.PHONY: dev-setup dev-start dev-stop dev-status

# Setup development environment
dev-setup: deps
	@echo "Setting up development environment..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d postgres pgadmin
	@echo "Waiting for database to be ready..."
	@sleep 10
	@echo "Development environment ready!"

# Start development environment
dev-start: dev-setup
	@echo "Starting development environment..."
	@echo "Database: localhost:5435"
	@echo "pgAdmin: http://localhost:5050"
	@echo "App: http://localhost:8080"

# Stop development environment
dev-stop:
	@echo "Stopping development environment..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down

# Check status of services
dev-status:
	@echo "Checking service status..."
	docker compose -f $(DOCKER_COMPOSE_FILE) ps

# Database commands
.PHONY: db-migrate db-migrate-up db-migrate-reset db-status db-psql db-create db-drop db-reset

# Run database migrations
db-migrate: db-migrate-up

# Apply all migrations
db-migrate-up:
	@echo "Applying database migrations..."
	@bash scripts/apply_migrations.sh

# Reset and reapply all migrations
db-migrate-reset:
	@echo "Resetting and reapplying migrations..."
	@echo "WARNING: This will drop all tables!"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		docker compose -f $(DOCKER_COMPOSE_FILE) down -v; \
		docker compose -f $(DOCKER_COMPOSE_FILE) up -d postgres; \
		echo "Waiting for database..."; \
		sleep 5; \
		bash scripts/apply_migrations.sh; \
	fi

# Show database status
db-status:
	@echo "Database status:"
	@docker compose -f $(DOCKER_COMPOSE_FILE) ps postgres

# Connect to database with psql
db-psql:
	@echo "Connecting to database..."
	@docker compose -f $(DOCKER_COMPOSE_FILE) exec postgres psql -U postgres -d jobot

# Create database
db-create:
	@echo "Creating database..."
	@docker compose -f $(DOCKER_COMPOSE_FILE) exec postgres psql -U postgres -c "CREATE DATABASE jobot;"

# Drop database
db-drop:
	@echo "Dropping database..."
	@echo "WARNING: This will delete all data!"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		docker compose -f $(DOCKER_COMPOSE_FILE) exec postgres psql -U postgres -c "DROP DATABASE IF EXISTS jobot;"; \
	fi

# Reset database (drop, create, migrate)
db-reset:
	@echo "Resetting database..."
	@echo "WARNING: This will delete all data!"
	@read -p "Are you sure? [y/N] " -n 1 -r; \
	echo; \
	if [[ $$REPLY =~ ^[Yy]$$ ]]; then \
		$(MAKE) db-drop; \
		$(MAKE) db-create; \
		$(MAKE) db-migrate; \
	fi

# Help
.PHONY: help
help:
	@echo "Available commands:"
	@echo ""
	@echo "Go commands:"
	@echo "  build     - Build the application"
	@echo "  run       - Run the application locally"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean build artifacts"
	@echo "  deps      - Install dependencies"
	@echo "  fmt       - Format code"
	@echo "  lint      - Lint code"
	@echo ""
	@echo "Docker commands:"
	@echo "  docker-build        - Build Docker image"
	@echo "  docker-up           - Start services with Docker Compose"
	@echo "  docker-down         - Stop Docker Compose services"
	@echo "  docker-down-volumes - Stop services and remove volumes"
	@echo "  docker-logs         - View all logs"
	@echo "  docker-logs-app     - View app logs"
	@echo "  docker-logs-db      - View database logs"
	@echo "  docker-rebuild      - Rebuild and restart services"
	@echo "  docker-clean        - Clean Docker resources"
	@echo ""
	@echo "API Documentation commands:"
	@echo "  api-docs         - Show API documentation URLs"
	@echo "  api-docs-open    - Open Swagger UI in browser"
	@echo "  api-validate     - Validate OpenAPI specification"
	@echo ""
	@echo "Development commands:"
	@echo "  dev-setup   - Setup development environment"
	@echo "  dev-start   - Start development environment"
	@echo "  dev-stop    - Stop development environment"
	@echo "  dev-status  - Check service status"
	@echo ""
	@echo "Database commands:"
	@echo "  db-migrate        - Apply all database migrations"
	@echo "  db-migrate-up     - Apply all database migrations"
	@echo "  db-migrate-reset  - Reset and reapply all migrations"
	@echo "  db-status         - Show database status"
	@echo "  db-psql           - Connect to database with psql"
	@echo "  db-create         - Create database"
	@echo "  db-drop           - Drop database"
	@echo "  db-reset          - Reset database (drop, create, migrate)"
	@echo ""
	@echo "Other:"
	@echo "  help        - Show this help message"
