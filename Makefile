.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build docker image to deploy
	docker build -t budougumi0617/gotodo:${DOCKER_TAG} \
		--target deploy ./

build-local: ## Build docker image to local development
	docker compose build --no-cache

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

test: ## Execute tests
	go test -race -shuffle=on ./...

backend-bash:
	docker compose exec backend bash

backend-faker-users:
	docker compose exec backend bash -c "go run src/commands/populateUsers.go"

backend-faker-products:
	docker compose exec backend bash -c "go run src/commands/populateProducts.go"

backend-faker-orders:
	docker compose exec backend bash -c "go run src/commands/populateOrders.go"

backend-redis-ranking:
	docker compose exec backend bash -c "go run src/commands/updateRankings.go"

dry-migrate: ## Try migration
	mysqldef -u gof1ber -p gof1ber -h 127.0.0.1 -P 33306 shop_site --dry-run < ./_tools/mysql/schema.sql

migrate:  ## Execute migration
	mysqldef -u gof1ber -p gof1ber -h 127.0.0.1 -P 33306 shop_site < ./_tools/mysql/schema.sql

generate: ## Generate codes
	go generate ./...

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'