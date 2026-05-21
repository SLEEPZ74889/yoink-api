.PHONY: run build test migrate

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