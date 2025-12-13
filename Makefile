.PHONY: all build run test clean tidy

APP_NAME=daily-uwoh
CMD_PATH=./cmd/server

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
