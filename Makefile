PROJECT_DIR := .
DEPLOYMENT_DIR := $(PROJECT_DIR)/deployment
DEV_COMPOSE := $(DEPLOYMENT_DIR)/dev/docker-compose.yml
PROD_COMPOSE := $(DEPLOYMENT_DIR)/prod/docker-compose.yml
SQLC_CONFIG := $(PROJECT_DIR)/internal/adapters/postgres/sqlc/sqlc.yaml

.PHONY: help server client build-server build-client test clean dev dev-build dev-build-d prod prod-build prod-build-d sqlc

## 🖥️ Start server application (Go run)
server:
	@echo "[i] Starting server with Go..."
	@go run ./cmd/server

## 🖥️ Start client application (Go run)
client:
	@echo "[i] Starting client with Go..."
	@go run ./cmd/client

## 🖥️ Start client application (Go run)
jwt:
	@echo "[i] Generate Jwt with Go..."
	@go run ./cmd/jwt

## 🧪 Run all tests recursively
test:
	@echo "[i] Running all tests..."
	@go test ./...

## 🧹 Clean built binaries and storage
clean:
	@echo "[i] Cleaning binaries and storage..."
	@rm -rf bin/*
	@rm -rf ./storage
	@rm -rf ./tmp

## 🐳 Start project in development mode
dev:
	@echo "[i] Starting dev environment..."
	@docker compose -f $(DEV_COMPOSE) up

## 🐳 Build and start dev environment
dev-build:
	@echo "[i] Building and starting dev environment..."
	@docker compose -f $(DEV_COMPOSE) up --build

## 🐳 Build and start dev environment (detached)
dev-build-d:
	@echo "[i] Building and starting dev environment in background..."
	@docker compose -f $(DEV_COMPOSE) up --build -d

## 🐳 Start project in production mode
prod:
	@echo "[i] Starting prod environment..."
	@docker compose -f $(PROD_COMPOSE) up

## 🐳 Build and start prod environment
prod-build:
	@echo "[i] Building and starting prod environment..."
	@docker compose -f $(PROD_COMPOSE) up --build

## 🐳 Build and start prod environment (detached)
prod-build-d:
	@echo "[i] Building and starting prod environment in background..."
	@docker compose -f $(PROD_COMPOSE) up --build -d

## 🧬 Generate SQLC code
sqlc:
	@echo "[i] Generating SQLC code..."
	@sqlc generate -f $(SQLC_CONFIG)

## 📋 Show available targets
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  server           Run server app (Go)"
	@echo "  client           Run client app (Go)"
	@echo "  test             Run all tests recursively"
	@echo "  clean            Remove built binaries and ./storage"
	@echo "  dev              Start dev docker-compose"
	@echo "  dev-build        Build and start dev environment"
	@echo "  dev-build-d      Build and start dev environment (detached)"
	@echo "  prod             Start prod docker-compose"
	@echo "  prod-build       Build and start prod environment"
	@echo "  prod-build-d     Build and start prod environment (detached)"
	@echo "  sqlc             Generate SQLC code"
	@echo "  help             Show this help"
