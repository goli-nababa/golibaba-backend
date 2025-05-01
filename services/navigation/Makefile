# Build directories
ROOT_DIR := $(shell pwd)
BUILD_DIR := $(ROOT_DIR)/build
POSTGRES_DIR := $(BUILD_DIR)/postgres
PROJECT_DIR := $(BUILD_DIR)/project
REDIS_DIR := $(BUILD_DIR)/redis

# Docker Compose Files
POSTGRES_COMPOSE := docker-compose.yaml
PROJECT_COMPOSE := docker-compose.yaml
REDIS_COMPOSE := docker-compose.yaml

# Docker network
NETWORK_NAME := golibaba-network

# Environment files
ENV_FILES := \
	$(POSTGRES_DIR)/.env \
	$(PROJECT_DIR)/.env \
	$(REDIS_DIR)/.env \
	$(PROJECT_DIR)/config/config.json

# Colors for pretty output
GREEN := "\033[0;32m"
RED := "\033[0;31m"
YELLOW := "\033[1;33m"
NC := "\033[0m"

# Default target
.DEFAULT_GOAL := help

# Check for required commands
check-prereqs:
	@echo ${YELLOW}"Checking prerequisites..."${NC}
	@which docker >/dev/null 2>&1 || (echo "docker is required but not installed. Aborting." && exit 1)
	@which docker-compose >/dev/null 2>&1 || (echo "docker-compose is required but not installed. Aborting." && exit 1)
	@which go >/dev/null 2>&1 || (echo "go is required but not installed. Aborting." && exit 1)
	@echo ${GREEN}"Prerequisites OK"${NC}

# Setup vendor directory
setup-deps: check-prereqs
	@echo ${YELLOW}"Setting up dependencies..."${NC}
	@if [ ! -d "vendor" ]; then \
		echo "Creating vendor directory..." && \
		go mod tidy && \
		go mod vendor; \
	fi
	@echo ${GREEN}"Dependencies ready"${NC}

# Create necessary networks
create-network:
	@echo ${YELLOW}"Creating docker networks..."${NC}
	@docker network create $(NETWORK_NAME) 2>/dev/null || true
	@echo ${GREEN}"Network ready"${NC}

# Create environment files from samples
create-env-files:
	@echo ${YELLOW}"Creating environment files..."${NC}
	@for dir in $(POSTGRES_DIR) $(PROJECT_DIR) $(REDIS_DIR); do \
		if [ ! -f $$dir/.env ] && [ -f $$dir/.env.sample ]; then \
			echo "Creating .env from sample in $$dir"; \
			cp $$dir/.env.sample $$dir/.env; \
		fi; \
	done
	@if [ ! -f $(PROJECT_DIR)/config/config.json ] && [ -f $(PROJECT_DIR)/config/config.sample.json ]; then \
		echo "Creating config.json from sample"; \
		mkdir -p $(PROJECT_DIR)/config && \
		cp $(PROJECT_DIR)/config/config.sample.json $(PROJECT_DIR)/config/config.json; \
	fi
	@echo ${GREEN}"Environment files ready"${NC}

# Initialize project
init: check-prereqs setup-deps create-network create-env-files
	@echo ${GREEN}"Project initialized"${NC}

# Start individual services
postgres-up:
	@echo ${YELLOW}"Starting PostgreSQL..."${NC}
	@cd $(POSTGRES_DIR) && docker-compose up -d
	@echo ${GREEN}"PostgreSQL is running"${NC}

redis-up:
	@echo ${YELLOW}"Starting Redis..."${NC}
	@cd $(REDIS_DIR) && docker-compose up -d
	@echo ${GREEN}"Redis is running"${NC}

app-up:
	@echo ${YELLOW}"Starting Application..."${NC}
	@cd $(PROJECT_DIR) && docker-compose up -d
	@echo ${GREEN}"Application is running"${NC}

# Main commands
up: init postgres-up redis-up app-up ## Start all services
	@echo ${GREEN}"All services are running"${NC}
	@make ps

build: init ## Build all services
	@echo ${YELLOW}"Building services..."${NC}
	@cd $(POSTGRES_DIR) && docker-compose build
	@cd $(REDIS_DIR) && docker-compose build
	@cd $(PROJECT_DIR) && docker-compose build
	@echo ${GREEN}"Build complete"${NC}

down: ## Stop all services
	@echo ${YELLOW}"Stopping services..."${NC}
	@cd $(PROJECT_DIR) && docker-compose down --remove-orphans
	@cd $(REDIS_DIR) && docker-compose down --remove-orphans
	@cd $(POSTGRES_DIR) && docker-compose down --remove-orphans
	@echo ${GREEN}"All services stopped"${NC}

ps: ## Show running services
	@echo ${YELLOW}"PostgreSQL containers:"${NC}
	@cd $(POSTGRES_DIR) && docker-compose ps
	@echo ${YELLOW}"Redis containers:"${NC}
	@cd $(REDIS_DIR) && docker-compose ps
	@echo ${YELLOW}"Application containers:"${NC}
	@cd $(PROJECT_DIR) && docker-compose ps

logs: ## Show logs from all services
	@cd $(POSTGRES_DIR) && docker-compose logs -f &
	@cd $(REDIS_DIR) && docker-compose logs -f &
	@cd $(PROJECT_DIR) && docker-compose logs -f

logs-app: ## Show only application logs
	@cd $(PROJECT_DIR) && docker-compose logs -f

logs-db: ## Show only database logs
	@cd $(POSTGRES_DIR) && docker-compose logs -f

logs-redis: ## Show only redis logs
	@cd $(REDIS_DIR) && docker-compose logs -f

clean: down ## Clean up everything
	@echo ${YELLOW}"Cleaning up..."${NC}
	@cd $(POSTGRES_DIR) && docker-compose down -v --remove-orphans
	@cd $(REDIS_DIR) && docker-compose down -v --remove-orphans
	@cd $(PROJECT_DIR) && docker-compose down -v --remove-orphans
	@docker system prune -f
	@rm -rf vendor
	@echo ${GREEN}"Cleanup complete"${NC}

restart: down up ## Restart all services

test: setup-deps ## Run tests
	@echo ${YELLOW}"Running tests..."${NC}
	@go test -v ./...
	@echo ${GREEN}"Tests complete"${NC}

help: ## Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*##"; printf "\033[36m"} \
		/^[a-zA-Z_-]+:.*?##/ { printf "  %-20s %s\n", $$1, $$2 } \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

.PHONY: help init build up down ps logs clean restart test \
	check-prereqs setup-deps create-network create-env-files \
	postgres-up redis-up app-up logs-app logs-db logs-redis