DOCKER_COMPOSE_DEV=compose.dev.yml

dev/up:
	docker compose -f $(DOCKER_COMPOSE_DEV) up -d
dev/down:
	docker compose -f $(DOCKER_COMPOSE_DEV) down --remove-orphans
dev/build:
	docker compose -f $(DOCKER_COMPOSE_DEV) build
dev/build/nocache:
	docker compose -f $(DOCKER_COMPOSE_DEV) build --no-cache
