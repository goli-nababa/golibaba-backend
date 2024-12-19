ROOT_DIR := ./

# Compose files relative to the root directory
COMPOSE_FILES := \
	-f $(ROOT_DIR)/compose/rabbitmq/docker-compose.yaml

NETWORK_NAME=golibaba-network

ensure-network:
	@echo "Ensuring project network is exists..."
	@if [ -z "$$(docker network ls --filter name=$(NETWORK_NAME) --format '{{.Name}}')" ]; then \
		echo "Network $(NETWORK_NAME) does not exist. Creating..."; \
		docker network create $(NETWORK_NAME); \
	fi

up: ensure-network
	docker compose --project-directory $(ROOT_DIR) $(COMPOSE_FILES) up