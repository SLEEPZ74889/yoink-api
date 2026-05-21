include .env
export

.PHONY: run build test migrate fmt

# Server starten
run:
	go run cmd/api/main.go

# Binary bauen
build:
	go build -o bin/yoink cmd/api/main.go

# Tests ausführen
test:
	go test ./...

# Code formatieren
fmt:
	go fmt ./...

# Datenbankmigration ausführen
migrate:
	psql "$(DATABASE_URL)" -f internal/db/migrations/001_init.sql