.PHONY: all build run test clean tidy

APP_NAME=daily-uwoh
CMD_PATH=./cmd/server
DB_URL=postgresql://postgres:shiratama@localhost:5432/dailyuwoh?sslmode=disable

all: build

build:
	@echo "Building $(APP_NAME)..."
	@go build -o bin/server.exe $(CMD_PATH)

run:
	@echo "Running $(APP_NAME)..."
	@go run $(CMD_PATH)

test:
	@echo "Running tests..."
	@go test ./... -v

clean:
	@echo "Cleaning up..."
	@if exist bin rmdir /s /q bin

tidy:
	@echo "Tidying modules..."
	@go mod tidy

sqlc:
	sqlc generate

createdb:
	docker exec -it postgres18 createdb --username=postgres --owner=postgres dailyuwoh

dropdb:
	docker exec -it postgres18 dropdb dailyuwoh

migrateup:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

server:
	go run $(CMD_PATH)

.PHONY: sqlc migrateup migratedown createdb dropdb
