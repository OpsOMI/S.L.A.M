DEV_COMPOSE=deployment/dev/docker-compose.yml
PROD_COMPOSE=deployment/prod/docker-compose.yml

.PHONY: help server client build-server build-client test clean dev-build dev-build-d prod-build prod-build-d

## 🖥️ Start server application (Go run)
server:
	@echo "[i] Starting server with Go..."
	@go run ./cmd/server

## 🖥️ Start client application (Go run)
client:
	@echo "[i] Starting client with Go..."
	@go run ./cmd/client

## 🛠️ Build server binary
build-server:
	@echo "[i] Building server binary..."
	@go build -o bin/server ./cmd/server

## 🛠️ Build client binary
build-client:
	@echo "[i] Building client binary..."
	@go build -o bin/client ./cmd/client

## 🧪 Run all tests recursively
test:
	@echo "[i] Running all tests..."
	@go test ./...

## 🧹 Clean built binaries and storage
clean:
	@echo "[i] Cleaning binaries and storage..."
	@rm -rf bin/*
	@rm -rf ./storage

## 🐳 Start development environment (foreground)
dev-build:
	@echo "[i] Starting dev docker-compose (foreground)..."
	@docker-compose -f $(DEV_COMPOSE) up --build

## 🐳 Start development environment (background)
dev-build-d:
	@echo "[i] Starting dev docker-compose (detached)..."
	@docker-compose -f $(DEV_COMPOSE) up --build -d

## 🐳 Start production environment (foreground)
prod-build:
	@echo "[i] Starting prod docker-compose (foreground)..."
	@docker-compose -f $(PROD_COMPOSE) up --build

## 🐳 Start production environment (background)
prod-build-d:
	@echo "[i] Starting prod docker-compose (detached)..."
	@docker-compose -f $(PROD_COMPOSE) up --build -d

## 📋 Show available targets
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  server           Run server app (Go)"
	@echo "  client           Run client app (Go)"
	@echo "  build-server     Build server binary"
	@echo "  build-client     Build client binary"
	@echo "  test             Run all tests"
	@echo "  clean            Remove built binaries and ./storage"
	@echo "  dev-build        Start dev docker-compose"
	@echo "  dev-build-d      Start dev docker-compose (detached)"
	@echo "  prod-build       Start prod docker-compose"
	@echo "  prod-build-d     Start prod docker-compose (detached)"
	@echo "  help             Show this help"
