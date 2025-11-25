.PHONY: all run test lint clean up-infra

all: lint test

up-infra:
	@echo "Starting Database and Cache..."
	docker-compose up -d db cache

lint:
	@echo "Running GolangCI-Lint..."
	golangci-lint run

test: up-infra
	@echo "Running Tests..."
	DB_HOST=localhost go test -v ./...

clean:
	@echo "Cleaning up..."
	docker compose down -v

run: clean
	@echo "Starting Pokedex API..."
	docker compose up --build -d

