go-mod-vendor:
	@echo "Running 'go mod vendor' to sync dependencies..."
	@go mod tidy
	@go mod vendor

ROOT_DIR := ./build

COMPOSE_FILES := \
	-f $(ROOT_DIR)/redis/docker-compose.yaml \
	-f $(ROOT_DIR)/project/docker-compose.yaml

up: go-mod-vendor
	@docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) up -d

build: go-mod-vendor
	@docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) build $(FLAGS) $(CONTAINER)

rebuild: go-mod-vendor
	@docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) up --build

down: go-mod-vendor
	docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) down