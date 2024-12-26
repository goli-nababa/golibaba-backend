NETWORK_NAME=golibaba-network

ensure-network:
	@echo "Ensuring project network is exists..."
	@if [ -z "$$(docker network ls --filter name=$(NETWORK_NAME) --format '{{.Name}}')" ]; then \
		echo "Network $(NETWORK_NAME) does not exist. Creating..."; \
		docker network create $(NETWORK_NAME); \
	fi

SERVICES := api_gateway $(wildcard services/*)

up: ensure-network
	@for service in $(SERVICES); do \
		$(MAKE) -C $$service up; \
	done


down: ensure-network
	@for service in $(shell echo $(SERVICES) | tr ' ' '\n' | tac); do \
		$(MAKE) -C $$service down; \
	done


