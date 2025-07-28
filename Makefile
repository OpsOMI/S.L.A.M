.PHONY: help server client build-server build-client test clean

## 🖥️ Start server application
server:
	@echo "[i] Starting server..."
	@go run ./cmd/server

## 🖥️ Start client application
client:
	@echo "[i] Starting client..."
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

## 🧹 Clean built binaries
clean:
	@echo "[i] Cleaning binaries..."
	@rm -rf bin/*

## 📋 Show available targets
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  server         Run server app"
	@echo "  client         Run client app"
	@echo "  build-server   Build server binary"
	@echo "  build-client   Build client binary"
	@echo "  test           Run all tests"
	@echo "  clean          Remove built binaries"
	@echo "  help           Show this help"
