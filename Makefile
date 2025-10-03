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
.PHONY: db-migrate db-seed db-reset

# Run database migrations
db-migrate:
	@echo "Running database migrations..."
	# Add migration commands here

# Seed database with sample data
db-seed:
	@echo "Seeding database..."
	# Add seed commands here

# Reset database
db-reset:
	@echo "Resetting database..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down -v
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d postgres

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
	@echo "Development commands:"
	@echo "  dev-setup   - Setup development environment"
	@echo "  dev-start   - Start development environment"
	@echo "  dev-stop    - Stop development environment"
	@echo "  dev-status  - Check service status"
	@echo ""
	@echo "Database commands:"
	@echo "  db-migrate  - Run database migrations"
	@echo "  db-seed     - Seed database with sample data"
	@echo "  db-reset    - Reset database"
	@echo ""
	@echo "Other:"
	@echo "  help        - Show this help message"
