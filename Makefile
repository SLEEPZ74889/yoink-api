include .env
export

.PHONY: run build test migrate fmt tidy

run:
	go run cmd/api/main.go

build:
	go build -o bin/yoink cmd/api/main.go

test:
	go test ./... -v

fmt:
	go fmt ./...

tidy:
	go mod tidy

migrate:
	psql "$(DATABASE_URL)" -f internal/db/migrations/001_init.sql
