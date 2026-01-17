# Load environment variables from .env file
include .env
export

# Database URL for migrations
DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_PATH=./migrations

# ==================== Migration Commands ====================

## migrate-up: Run all pending migrations
migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" up

## migrate-down: Rollback all migrations
migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down

## migrate-down-1: Rollback last migration
migrate-down-1:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" down 1

## migrate-version: Show current migration version
migrate-version:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" version

## migrate-force VERSION=x: Force set migration version (use when dirty)
migrate-force:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" force $(VERSION)

## migrate-create NAME=xxx: Create a new migration file
migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(NAME)

## migrate-drop: Drop everything in the database
migrate-drop:
	migrate -path $(MIGRATIONS_PATH) -database "$(DATABASE_URL)" drop -f

# ==================== App Commands ====================

## run: Run the application
run:
	go run cmd/main.go

## build: Build the application
build:
	go build -o bin/app cmd/main.go

## test: Run tests
test:
	go test -v ./...

## swagger: Generate swagger docs
swagger:
	swag init -g cmd/main.go -o docs

# ==================== Docker Commands ====================

## docker-up: Start docker containers
docker-up:
	docker-compose up -d

## docker-down: Stop docker containers
docker-down:
	docker-compose down

## docker-logs: View docker logs
docker-logs:
	docker-compose logs -f

# ==================== Help ====================

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

.PHONY: migrate-up migrate-down migrate-down-1 migrate-version migrate-force migrate-create migrate-drop run build test swagger docker-up docker-down docker-logs help
.DEFAULT_GOAL := help
